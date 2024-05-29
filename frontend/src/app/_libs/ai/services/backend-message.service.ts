import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ChatMessage } from '../../models/chat-message.model';
import { ConfigService } from './config.service';

@Injectable({
    providedIn: 'root',
})
export class BackendMessageService {
    constructor(
        private _http: HttpClient,
        private _config: ConfigService,
    ) {}

    public create(chatMessage: ChatMessage): Observable<ChatMessage> {
        return this._http.post<{ status: number; data: ChatMessage }>(this._config.backendUrlPersistence + 'messages', chatMessage).pipe(
            map((res) => {
                return res.data;
            }),
        );
    }

    public getByChatId(chatId: string): Observable<ChatMessage[]> {
        return this._http.get<ChatMessage[]>(this._config.backendUrlPersistence + 'messages/odata?$filter=chatId eq ' + chatId);
    }
}
