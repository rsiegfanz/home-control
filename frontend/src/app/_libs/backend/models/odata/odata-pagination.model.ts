import { IOdata } from './odata.interface';
import { IPaginationDto } from '../../dtos/pagination-dto.interface';

export class ODataPagination implements IOdata {
    public pageSize = 15;

    public totalPages = 0;

    public totalElements = 0;

    public page = 1;

    public toGetParameter(): string {
        let odataString = `$top=${this.pageSize}`;
        odataString += `&$skip=${(this.page - 1) * this.pageSize}`;
        return odataString;
    }

    public static fromDto(dto: IPaginationDto): ODataPagination {
        const obj = new ODataPagination();
        obj.pageSize = dto.pageSize;
        obj.totalPages = dto.totalPages;
        obj.totalElements = dto.totalElements;
        obj.page = dto.page;
        return obj;
    }
}
