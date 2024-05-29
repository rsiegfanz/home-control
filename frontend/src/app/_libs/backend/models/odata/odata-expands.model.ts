import { IOdata } from './odata.interface';

export class OdataExpands implements IOdata {
    public expands: string[] = [];

    public constructor(expands: string[] = []) {
        this.expands = expands;
    }

    public addField(field: string): void {
        this.expands.push(field);
    }

    public toGetParameter(): string {
        return this.expands.length > 0 ? `$expand=${this.expands.join(',')}` : '';
    }
}
