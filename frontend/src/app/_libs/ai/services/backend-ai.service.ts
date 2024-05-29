import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, map, catchError } from 'rxjs';
import { ChatMessage } from '../models/chat-message.model';
import { ConfigService } from './config.service';

@Injectable({
    providedIn: 'root',
})
export class BackendAiService {
    constructor(
        private _http: HttpClient,
        private _config: ConfigService,
    ) {}

    public stt(data: Blob): Observable<string> {
        const formData = new FormData();
        formData.append('file', data);
        return this._http.post(this._config.backendUrlStt, formData).pipe(
            map((res: any) => {
                return res.text as string;
            }),
            catchError((err) => {
                console.log(err);
                return 'test';
            }),
        );
    }

    public tts(message: string): Observable<Blob> {
        return this._http.post(this._config.backendUrlTts, { text: message }, { observe: 'response', responseType: 'blob' }).pipe(
            map((res: any) => {
                // console.log('result is ', res);
                // console.log(res.body);
                return res.body;
            }),
        );
    }

    public chat(chatId: string): Observable<ChatMessage> {
        return this._http.post(this._config.backendUrlChat, { chatId }).pipe(
            map((res: any) => {
                return res as ChatMessage;
            }),
        );
    }
}
