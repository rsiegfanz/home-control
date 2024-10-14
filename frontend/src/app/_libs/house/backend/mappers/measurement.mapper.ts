import { DateTime } from 'luxon';
import Measurement from '../../models/measurement.model';
import { IMeasurementDto, IClimateMeasurementSchemaDto } from '../dtos/measurement.dto';

export function mapDtoToModel(dto: IMeasurementDto | undefined): Measurement | undefined {
    if (!dto) {
        return undefined;
    }

    return new Measurement(dto.timestamp, dto.temperature, dto.humidity);
}

export function mapGraphlQLDtoToModel(dto: IClimateMeasurementSchemaDto): Measurement {
    const parsedDate = DateTime.fromFormat(dto.recordedAt.replace('CEST', '').trimEnd(), 'yyyy-MM-dd HH:mm:ss ZZZ', { zone: 'local' });
    return new Measurement(new Date(parsedDate.toJSDate()), dto.temperature, dto.humidity);
}
