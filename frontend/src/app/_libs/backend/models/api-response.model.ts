import { BaseModel } from './base.model';
import { OdataResponse } from './odata/odata-response.interface';

export class ApiResponse<TModel extends BaseModel> {
    public status!: number;

    public data: TModel | undefined;

    public odata: OdataResponse<TModel> | undefined;

    public messages: string[] = [];

    public isError: boolean = false;
}
