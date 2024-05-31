import { IBaseRequestDto } from '../../../backend/dtos/base-request-dto.interface';
import { IBaseResponseDto } from '../../../backend/dtos/base-response-dto.interface';

export interface MeasurementRequestDto extends IBaseRequestDto {}

export interface MeasurementResponseDto extends IBaseResponseDto {
    timestamp: Date;
    temperature: number;
    humidity: number;
}
