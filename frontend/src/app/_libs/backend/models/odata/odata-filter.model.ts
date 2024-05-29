import { IOdata } from './odata.interface';
import { EFilterOperator, EFilterTypes } from './filter/filter.enums';

export class OdataFilter implements IOdata {
    public constructor(
        public field?: string,
        public values: string[] | number[] = [0, 0],
        public operator?: EFilterOperator,
        public type?: EFilterTypes,
    ) {}

    public toGetParameter(): string {
        if (!this.field || !this.operator || !this.type) {
            return '';
        }
        switch (this.operator) {
            case EFilterOperator.BETWEEN:
                if (this.type === EFilterTypes.PACE) {
                    this.values.sort((a, b) => (a > b ? -1 : 1));
                }
                return `${this.field} ge ${this._getTypeDependentValue(0)} and ${this.field} le ${this._getTypeDependentValue(1)}`;
            case EFilterOperator.CONTAINS:
                return `contains(${this.field},'${this.values[0]}')`;
            default:
                return `${this.field} ${this._getODataOperator()} ${this._getTypeDependentValue()}`;
        }
    }

    public availableOperators(): EFilterOperator[] {
        switch (this.type) {
            case EFilterTypes.NUMERIC:
                return [EFilterOperator.GREATER, EFilterOperator.GREATEREQUALS, EFilterOperator.BETWEEN, EFilterOperator.LESS, EFilterOperator.LESSEQUALS, EFilterOperator.EQUALS];
            case EFilterTypes.STRING:
                return [EFilterOperator.CONTAINS];
            case EFilterTypes.DATE:
                return [EFilterOperator.GREATER, EFilterOperator.GREATEREQUALS, EFilterOperator.BETWEEN, EFilterOperator.LESS, EFilterOperator.LESSEQUALS, EFilterOperator.EQUALS];
            case EFilterTypes.TIME:
            case EFilterTypes.PACE:
                return [EFilterOperator.GREATER, EFilterOperator.GREATEREQUALS, EFilterOperator.BETWEEN, EFilterOperator.LESS, EFilterOperator.LESSEQUALS, EFilterOperator.EQUALS];
        }
        return [EFilterOperator.GREATER, EFilterOperator.GREATEREQUALS, EFilterOperator.BETWEEN, EFilterOperator.LESS, EFilterOperator.LESSEQUALS, EFilterOperator.EQUALS];
    }

    private _getTypeDependentValue(index = 0): string | number {
        if (this.type === EFilterTypes.STRING) {
            return `'${this.values[index]}'`;
        }

        if (this.type === EFilterTypes.TIME) {
            if (!(this.values[index] as string).includes(':')) {
                console.warn(this.values[index], 'not a valid time');
                return 0;
            }
            const minutesSeconds = (this.values[index] as string).split(':');
            return parseInt(minutesSeconds[0], 10) * 60 + parseInt(minutesSeconds[1], 10);
        }

        return this.values[index];
    }

    private _getODataOperator(): string {
        switch (this.operator) {
            case EFilterOperator.EQUALS:
                return 'eq';
            case EFilterOperator.LESS:
                return 'lt';
            case EFilterOperator.LESSEQUALS:
                return 'le';
            case EFilterOperator.GREATER:
                return 'gt';
            case EFilterOperator.GREATEREQUALS:
                return 'ge';
            case EFilterOperator.BETWEEN:
                return 'between';
            case EFilterOperator.CONTAINS:
                return 'contains';
            default:
                return 'eq';
        }
    }
}
