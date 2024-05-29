import { IOdata } from './odata.interface';

export enum ESortDirection {
    ASC = 'asc',
    DESC = 'desc',
}

export class ODataOrder implements IOdata {
    public column = 'id';

    public direction = ESortDirection.ASC;

    public toGetParameter(): string {
        return `$orderby=${this.column} ${this.direction}`;
    }
}
