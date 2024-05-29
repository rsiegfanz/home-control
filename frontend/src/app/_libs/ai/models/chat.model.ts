import { ChatMessage } from './chat-message.model';

export interface Chat {
    id: string;
    createdAt: string;
    messages: ChatMessage[];
}
