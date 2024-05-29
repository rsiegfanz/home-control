import { IBaseResponseDto } from './base-response-dto.interface';

export interface IApiResponseDto<TBaseDto extends IBaseResponseDto> {
    messages: string[];
    data: TBaseDto | undefined;
    status: number;
}
