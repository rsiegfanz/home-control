import { OdataFilter } from './odata-filter.model';
import { IOdata } from './odata.interface';

export enum FilterConcatenation {
    OR = ' or ',
    AND = ' and ',
}

export class OdataFilterCollection implements IOdata {
    public andFilters: OdataFilter[] = [];
    // public orFilters: OdataFilter[] = [];

    public toGetParameter(): string {
        let odataString = '$filter=';
        const filters = this.andFilters.filter((x) => this._isFilterValid(x));
        if (filters.length === 0) {
            return '';
        }
        odataString += filters.map((x) => x.toGetParameter()).join(FilterConcatenation.AND);
        return odataString;
    }

    public addAnd(andFilter: OdataFilter): void {
        this.andFilters.push(andFilter);
    }

    public remove(field: string | undefined): void {
        const idx = this.andFilters.findIndex((f) => f.field === field);
        if (idx >= 0) {
            this.andFilters.splice(idx, 1);
        }
    }

    public hasValidFilters(): boolean {
        return this.andFilters.filter((f) => this._isFilterValid(f)).length > 0;
    }

    private _isFilterValid(filter: OdataFilter): boolean {
        return !(!filter.type || !filter.field || !filter.operator);
    }

    // public addOr(field: string, value: string | number | number[], filterType: EODataFilterTypes): void {
    //   const filter = OdataFilterFactory.createFilter(filterType, field, value);
    //   if (filter) {
    //     this.orFilters.push(filter);
    //   }
    // }
}
