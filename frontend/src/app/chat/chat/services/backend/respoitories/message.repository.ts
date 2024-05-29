import { Injectable } from '@angular/core';
import {BaseApiGatewayRepository} from "../../../../_shared/services/backend/repositories/base-api-gateway.repository";
import {Message} from "../../../models/message.model";
import {IMessageRequestDto, IMessageResponseDto} from "../../../dtos/message-dto.interface";
import {MessageMapper} from "../mappers/message.mapper";

@Injectable({
  providedIn: 'root'
})
export class MessageRepository extends BaseApiGatewayRepository<Message, IMessageResponseDto, IMessageRequestDto> {

  protected mapper = new MessageMapper();

  protected getUrlSegment(): string {
    return '/persistence/messages';
  }
}
