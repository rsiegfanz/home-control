import { IconDefinition } from '@fortawesome/free-solid-svg-icons';

export default class Room {
    public readonly id: number;

    public readonly name: string;

    public readonly icon: IconDefinition;

    constructor(id: number, name: string, icon: IconDefinition) {
        this.id = id;
        this.name = name;
        this.icon = icon;
    }
}
