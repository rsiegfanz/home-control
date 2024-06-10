import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { ApiResponse } from '../../../_libs/backend/models/api-response.model';
import { OData } from '../../../_libs/backend/models/odata/odata.model';
import { Chat } from '../../chat/models/chat.model';
import { CommonModule } from '@angular/common';
import { IconDataprovider } from '../../../_libs/icons/icon.dataprovider';
import { FaIconComponent } from '@fortawesome/angular-fontawesome';

@Component({
    selector: 'app-chat-sidebar',
    standalone: true,
    imports: [CommonModule, FaIconComponent],
    templateUrl: './chat-sidebar.component.html',
    styleUrl: './chat-sidebar.component.scss',
})
export class ChatSidebarComponent implements OnInit {
    @ViewChild('inputNewChat') inputNewChat!: ElementRef;

    public toggleInputNewChat = false;

    public iconProvider = IconDataprovider;

    public chats$!: Observable<ApiResponse<Chat>>;

    public constructor(
        //        private readonly _chatRepository: ChatRepository,
        private _router: Router,
    ) {}

    public ngOnInit(): void {
        // const odata = new OData();
        //this.chats$ = this._chatRepository.odata(odata).pipe(tap(console.log));
        // .pipe(
        // map((response: ApiResponse<Chat>) => {
        //   console.log(response)
        //   if (response.isError) {
        //     console.error(response);
        //     return [];
        //   }
        //   console.log(response.odata!.items)
        //   return response.odata!.items || [];
        // })
        // );
    }

    public addNewChat(): void {
        this.toggleInputNewChat = true;
        setTimeout(() => {
            // eslint-disable-next-line @typescript-eslint/no-unsafe-call, @typescript-eslint/no-unsafe-member-access
            this.inputNewChat.nativeElement.focus();
        }, 0);
        console.log('new chat!');
        //        this._chatRepository.insert(new Chat()).subscribe((response: ApiResponse<Chat>) => {
        //            if (response.isError) {
        //                console.error(response);
        //                return;
        //            }
        //            console.log(response.data);
        //            this._router.navigate(['/chat', response.data!.id]);
        //        });
    }

    public submitNewChat(): void {
        this.toggleInputNewChat = false;
        console.log('submit chat');
    }

    public cancelNewChat(): void {
        this.toggleInputNewChat = false;
        console.log('close new chat!');
    }

    public selectChat(): void {
        console.log('select chat!');
    }
}
