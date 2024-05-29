import { IPaginationDto } from './pagination-dto.interface';

export interface IApiODataResponseDto<TBaseDto> {
    status: number;
    data: {
        items: TBaseDto[];
        pagination: IPaginationDto;
    };
}
