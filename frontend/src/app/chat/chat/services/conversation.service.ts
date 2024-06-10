import { ChatRepository } from './backend/respoitories/chat.repository';
import { MessageRepository } from './backend/respoitories/message.repository';
import { Injectable, signal } from '@angular/core';
import { AiRepository } from './backend/respoitories/ai.repository';
import { Message } from '../models/message.model';
import { EMessageTypes } from '../enums/message-types.enum';
import { BehaviorSubject, catchError, finalize, map, Observable, of, switchMap, tap } from 'rxjs';
import { OData } from '../../_shared/models/odata/odata.model';
import { OdataFilterCollection } from '../../_shared/models/odata/odata-filter-collection.model';
import { OdataFilter } from '../../_shared/models/odata/odata-filter.model';
import { EFilterOperator, EFilterTypes } from '../../_shared/models/odata/filter/filter.enums';
import { ESortDirection, ODataOrder } from '../../_shared/models/odata/odata-order.model';

@Injectable({
    providedIn: 'root',
})
export class ConversationService {
    // public chatId = new BehaviorSubject<string | undefined>(undefined);

    //     public messages$ = signal<Message[]>([]);

    //    public loadingAi$ = new BehaviorSubject(false);

    constructor(
        private readonly _chatRepo: ChatRepository,
        private readonly _messageRepo: MessageRepository,
        private readonly _aiRepo: AiRepository,
    ) {
        /*        this.chatId
            .pipe(
                switchMap((chatId) => {
                    console.log('trigger message loading', chatId);
                    if (!chatId) {
                        this.messages$.set([]);
                        return of();
                    }
                    const odata = new OData();
                    odata.filter = new OdataFilterCollection();
                    odata.filter.addAnd(new OdataFilter('chatId', [chatId], EFilterOperator.EQUALS, EFilterTypes.STRING));
                    odata.order = new ODataOrder();
                    odata.order.column = 'createdAt';
                    odata.order.direction = ESortDirection.DESC;
                    return this._messageRepo.odata(odata).pipe(
                        tap(console.log),
                        map((apiResponse) => {
                            if (apiResponse.status === 200) {
                                apiResponse.odata!.items.sort((a: Message, b: Message) => new Date(a.createdAt!).getTime() - new Date(b.createdAt!).getTime());
                                this.messages$.set(apiResponse.odata!.items);
                                return;
                            }
                            console.error(apiResponse);
                            this.messages$.set([]);
                        }),
                    );
                }),
            )
            .subscribe();
            */
    }

    public createChat(name: string): Observable<Chat> {
        //return this._chatRepo.insert(new Chat());
    }

    public addDocumentToChat(document: File): Observable<void> {
        /*        const chatIdValue = this.chatId.getValue();
        if (!chatIdValue) {
            return of();
        }
        return this._aiRepo.rag(document, chatIdValue);
      */
    }

    public saveMessage(message: string, shouldAiAnswer = true): Observable<Message | undefined> {
        /*
        const chatId = this.chatId.getValue();
        if (!chatId) {
            throw new Error('No chatId set');
        }
        const messageObj = new Message();
        messageObj.message = message;
        messageObj.messageType = EMessageTypes.HUMAN;
        messageObj.chatId = chatId;
        const insertObs$ = this._insertMessage(messageObj);
        if (!shouldAiAnswer) {
            return insertObs$;
        }
        return insertObs$.pipe(switchMap(() => this.getAiAnswer()));
*/
    }

    public getAiAnswer(): Observable<Message | undefined> {
        /*        const chatId = this.chatId.getValue();
        if (!chatId) {
            throw new Error('No chatId set');
        }
        this.loadingAi$.next(true);
        return this._aiRepo.sendMessage(chatId).pipe(
            tap((aiMessage) => this.messages$.update((values) => [...values, aiMessage])),
            catchError((err) => {
                console.error(err);
                return of(undefined);
            }),
            finalize(() => this.loadingAi$.next(false)),
        );
*/
    }

    private _insertMessage(messageObj: Message) {
        /*        return this._messageRepo.insert(messageObj).pipe(
            map((apiResponse) => {
                this.messages$.update((values) => [...values, apiResponse.data!]);
                if (apiResponse.status === 201) {
                    console.log('message saved', apiResponse);
                    console.log(this.messages$());
                    return apiResponse.data!;
                }
                throw new Error('Failed to save message');
            }),
            catchError((err) => {
                console.error(err);
                return of(undefined);
            }),
        );
*/
    }
}
