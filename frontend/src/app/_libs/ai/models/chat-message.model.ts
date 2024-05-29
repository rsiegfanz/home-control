export interface ChatMessage {
    status: string;
    message: string;
    messageType: string;
    createdAt: string;
    id: string;
    chatId: string;
    audioBlob: Blob | null;
}
