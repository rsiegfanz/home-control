import { BaseDbModel } from '../../../_libs/backend/models/base-db.model';
import { EMessageTypes } from '../enums/message-types.enum';

export class Message extends BaseDbModel {
    public createdAt: Date | undefined;

    public message!: string;

    public messageType!: EMessageTypes;

    public chatId!: string;
}
