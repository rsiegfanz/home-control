import { ODataPagination } from './odata-pagination.model';
import { BaseModel } from '../base.model';

export class OdataResponse<TModel extends BaseModel> {
    public items: TModel[] = [];

    public pagination!: ODataPagination;
}
