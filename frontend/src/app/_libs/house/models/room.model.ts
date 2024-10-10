import { IconDefinition } from '@fortawesome/free-solid-svg-icons';
import { BaseModel } from '../../backend/models/base.model';

export default class Room extends BaseModel {
    public readonly id: number;

    public readonly externalRoomId: string;

    public readonly name: string;

    public readonly icon: IconDefinition;

    constructor(id: number, externalId: string, name: string, icon: IconDefinition) {
        super();
        this.id = id;
        this.externalRoomId = externalId;
        this.name = name;
        this.icon = icon;
    }
}
