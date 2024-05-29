import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Message } from '../../../models/message.model';
import { environment } from '../../../../../environments/environment';
import { map, Observable, of } from 'rxjs';
import { IMessageResponseDto } from '../../../dtos/message-dto.interface';
import { mapResponseDtoToModel } from '../mappers/message.mapper';

@Injectable({
    providedIn: 'root',
})
export class AiRepository {
    constructor(private http: HttpClient) {}

    public rag(document: File, chatId: string): Observable<void> {
        const formData = new FormData();
        formData.append('file', document);
        formData.append('chatId', chatId);
        return this.http.post(environment.apiUrl + '/ai/rag/' + chatId, formData).pipe(
            map(() => {
                return;
            }),
        );
    }

    public sendMessage(chatId: string): Observable<Message> {
        return this.http.post<IMessageResponseDto>(environment.apiUrl + '/ai/message', { chatId }).pipe(
            map((res) => {
                return mapResponseDtoToModel(res);
            }),
        );
    }
}
