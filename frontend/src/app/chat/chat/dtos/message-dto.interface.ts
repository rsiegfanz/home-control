import {EMessageTypes} from "../enums/message-types.enum";

export interface IMessageRequestDto {
  id: string | undefined;
  message: string;
  messageType: EMessageTypes;
  chatId: string;
}

export interface IMessageResponseDto {
  id: string;
  message: string;
  messageType: EMessageTypes;
  chatId: string;
  createdAt: Date;
}
