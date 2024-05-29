import { OdataFilterCollection } from './odata-filter-collection.model';
import { ODataPagination } from './odata-pagination.model';
import { IOdata } from './odata.interface';
import { ODataOrder } from './odata-order.model';
import { OdataExpands } from './odata-expands.model';

export class OData implements IOdata {
    public pagination: ODataPagination = new ODataPagination();

    public filter = new OdataFilterCollection();

    public expands = new OdataExpands();

    public order = new ODataOrder();

    public toGetParameter(): string {
        let odataString = 'odata?' + this.pagination.toGetParameter();
        if (this.filter.hasValidFilters()) {
            odataString += '&' + this.filter.toGetParameter();
        }
        if (this.order) {
            odataString += '&' + this.order.toGetParameter();
        }
        if (this.expands) {
            odataString += '&' + this.expands.toGetParameter();
        }
        return odataString;
    }
}
