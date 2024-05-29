import { IBaseRequestDto } from '../../../backend/dtos/base-request-dto.interface';
import { IBaseResponseDto } from '../../../backend/dtos/base-response-dto.interface';

export interface TemperatureRequestDto extends IBaseRequestDto {}

export interface TemperatureResponseDto extends IBaseResponseDto {
    timestamp: Date;
    value: number;
}
