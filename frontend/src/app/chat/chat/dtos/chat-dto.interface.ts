import {EMessageTypes} from "../enums/message-types.enum";
import {IMessageResponseDto} from "./message-dto.interface";

export interface IChatRequestDto {
}

export interface IChatResponseDto {
  id: string;
  messages: IMessageResponseDto[];
  createdAt: Date;
}
