import { IChatRequestDto, IChatResponseDto } from '../../../dtos/chat-dto.interface';
import { Chat } from '../../../models/chat.model';
import { mapResponseDtoToModel as messageMapResponseDtoToModel } from './message.mapper';

export function mapModelToRequestDto(model: Chat): IChatRequestDto {
    return {};
}

export function mapResponseDtoToModel(dto: IChatResponseDto): Chat {
    const model = new Chat();
    model.id = dto.id;
    model.createdAt = new Date(dto.createdAt);
    model.messages = dto.messages ? dto.messages.map((message) => messageMapResponseDtoToModel(message)) : [];
    return model;
}
