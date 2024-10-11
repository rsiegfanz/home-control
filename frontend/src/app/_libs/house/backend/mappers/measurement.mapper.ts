import Measurement from '../../models/measurement.model';
import { IMeasurementDto, IClimateMeasurementSchemaDto } from '../dtos/measurement.dto';

export function mapDtoToModel(dto: IMeasurementDto | undefined): Measurement | undefined {
    if (!dto) {
        return undefined;
    }

    return new Measurement(dto.timestamp, dto.temperature, dto.humidity);
}

export function mapGraphlQLDtoToModel(dto: IClimateMeasurementSchemaDto): Measurement {
    return new Measurement(new Date(dto.recordedAt), dto.temperature, dto.humidity);
}
