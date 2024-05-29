import { Component } from '@angular/core';
import { ChatSidebarComponent } from './chat-sidebar/chat-sidebar.component';
import { ChatMessagesComponent } from './chat-messages/chat-messages.component';

@Component({
    selector: 'app-chat-home',
    standalone: true,
    imports: [ChatSidebarComponent, ChatMessagesComponent],
    templateUrl: './chat-home.component.html',
    styleUrl: './chat-home.component.scss',
})
export class ChatHomeComponent {}
