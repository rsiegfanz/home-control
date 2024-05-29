import Temperature from '../../models/temperature.model';
import { TemperatureRequestDto, TemperatureResponseDto } from '../dtos/temperature.dto';

export function mapResponseDtoToModel(dto: TemperatureResponseDto): Temperature {
    return new Temperature(dto.timestamp, dto.value);
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export function mapModelToRequestDto(model: Temperature): TemperatureRequestDto {
    throw Error('Not implemented');
}
