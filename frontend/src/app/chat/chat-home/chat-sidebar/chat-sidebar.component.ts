import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, tap } from 'rxjs';
import { ApiResponse } from '../../../_libs/backend/models/api-response.model';
import { OData } from '../../../_libs/backend/models/odata/odata.model';
import { Chat } from '../../chat/models/chat.model';
import { ChatRepository } from '../../chat/services/backend/respoitories/chat.repository';

@Component({
    selector: 'app-chat-sidebar',
    standalone: true,
    imports: [],
    templateUrl: './chat-sidebar.component.html',
    styleUrl: './chat-sidebar.component.scss',
})
export class ChatSidebarComponent implements OnInit {
    public chats$!: Observable<ApiResponse<Chat>>;

    public constructor(
        //        private readonly _chatRepository: ChatRepository,
        private _router: Router,
    ) {}

    public ngOnInit(): void {
        const odata = new OData();
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
        //        this._chatRepository.insert(new Chat()).subscribe((response: ApiResponse<Chat>) => {
        //            if (response.isError) {
        //                console.error(response);
        //                return;
        //            }
        //            console.log(response.data);
        //            this._router.navigate(['/chat', response.data!.id]);
        //        });
    }
}
