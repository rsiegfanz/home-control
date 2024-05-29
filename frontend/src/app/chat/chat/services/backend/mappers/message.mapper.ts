import { Message } from '../../../models/message.model';
import { IMessageRequestDto, IMessageResponseDto } from '../../../dtos/message-dto.interface';

export function mapModelToRequestDto(model: Message): IMessageRequestDto {
    return {
        id: model.id,
        message: model.message,
        messageType: model.messageType,
        chatId: model.chatId,
    };
}

export function mapResponseDtoToModel(dto: IMessageResponseDto): Message {
    const model = new Message();
    model.id = dto.id;
    model.message = dto.message;
    model.messageType = dto.messageType;
    model.chatId = dto.chatId;
    model.createdAt = dto.createdAt;
    return model;
}
