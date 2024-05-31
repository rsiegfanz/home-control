import Measurement from '../../models/measurement.model';
import { MeasurementRequestDto, MeasurementResponseDto } from '../dtos/measurement.dto';

export function mapResponseDtoToModel(dto: MeasurementResponseDto): Measurement {
    return new Measurement(dto.timestamp, dto.temperature, dto.humidity);
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export function mapModelToRequestDto(model: Measurement): MeasurementRequestDto {
    throw Error('Not implemented');
}
