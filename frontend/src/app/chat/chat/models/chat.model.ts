import { BaseDbModel } from '../../../_libs/backend/models/base-db.model';
import { Message } from './message.model';

export class Chat extends BaseDbModel {
    public createdAt: Date | undefined;

    public messages: Message[] = [];
}
