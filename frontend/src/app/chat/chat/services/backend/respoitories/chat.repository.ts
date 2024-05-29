import { Injectable } from '@angular/core';
import { BaseApiRepository } from '../../../../../_libs/backend/services/repositories/base-api.repository';
import { IChatRequestDto, IChatResponseDto } from '../../../dtos/chat-dto.interface';
import { Chat } from '../../../models/chat.model';
import { environment } from '../../../../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class ChatRepository extends BaseApiRepository<Chat, IChatResponseDto, IChatRequestDto> {
    protected override url = environment.backendAiUrl;

    protected override path = '';

    public override mapResponseDtoToModel(dto: IChatResponseDto): Chat {
        throw new Error('Method not implemented.');
    }

    public override mapModelToRequestDto(model: Chat): IChatRequestDto {
        throw new Error('Method not implemented.');
    }

    protected getUrlSegment(): string {
        return '/persistence/chats';
    }
}
