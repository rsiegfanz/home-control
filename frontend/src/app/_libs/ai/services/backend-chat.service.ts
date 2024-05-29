import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { Chat } from '../../models/chat.model';
import { ConfigService } from './config.service';

@Injectable({
    providedIn: 'root',
})
export class BackendChatService {
    constructor(
        private _http: HttpClient,
        private _config: ConfigService,
    ) {}

    public get(id: string, withMessages: boolean = true): Observable<Chat> {
        let url = this._config.backendUrlPersistence + 'chats/odata?$filter=id eq ' + id;
        if (withMessages) {
            url += '&$expand=messages';
        }
        return this._http.get<{ status: number; data: { items: Chat[] } }>(url).pipe(
            map((res) => {
                if (res.data.items.length === 0) {
                    throw new Error('Chat not found');
                }
                return res.data.items[0];
            }),
        );
    }

    public list(): Observable<Chat[]> {
        // const fakeResponse = [
        //   {
        //     id: '924994a9-c609-4c3c-9a4d-a7ea3279393d',
        //     createdAt: '2024-05-08T09:15:24.000Z',
        //   } as Chat,
        // ];
        // return of(fakeResponse);
        return this._http.get(`${this._config.backendUrlPersistence}chats/odata?$expand=messages`).pipe(
            map((res: any) => {
                if (res !== 0) {
                    return res.data.items as Chat[];
                }
                return [];
            }),
        );
    }

    public create(): Observable<Chat> {
        return this._http.post<{ status: number; data: Chat }>(this._config.backendUrlPersistence + 'chats', {}).pipe(
            map((res) => {
                return res.data;
            }),
        );
    }

    public delete(id: string): Observable<void> {
        return this._http.delete<void>(this._config.backendUrlPersistence + 'chats/' + id);
    }
}
