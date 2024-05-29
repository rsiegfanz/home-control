import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ChatMessage } from '../../models/chat-message.model';
import { Chat } from '../chat.model';
import { BackendMessageService } from './backend-message.service';

@Injectable({
    providedIn: 'root',
})
export class ChatService {
    public chat: Chat | undefined;

    public constructor(private readonly _persistenceService: BackendMessageService) {}

    public selectChat(chat: Chat) {
        this.chat = chat;
    }

    public addMessage(message: ChatMessage): Observable<ChatMessage> {
        if (!this.chat) {
            throw new Error('no parent chat');
        }

        if (!this.chat.messages) {
            this.chat.messages = [];
        }

        return this._persistenceService.create(message).pipe(
            map((res) => {
                this.chat!.messages.push(res);
                return res;
            }),
        );
    }
}
