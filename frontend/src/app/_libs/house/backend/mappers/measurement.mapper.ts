import Measurement from '../../models/measurement.model';
import { MeasurementDto } from '../dtos/measurement.dto';

export function mapDtoToModel(dto: MeasurementDto | undefined): Measurement | undefined {
    if (!dto) {
        return undefined;
    }

    return new Measurement(dto.timestamp, dto.temperature, dto.humidity);
}
