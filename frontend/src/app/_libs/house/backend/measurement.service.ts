import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { ApiResponse } from '../../backend/models/api-response.model';
import { BaseApiRepository } from '../../backend/services/repositories/base-api.repository';
import Measurement from '../models/measurement.model';
import { MeasurementRequestDto, MeasurementResponseDto } from './dtos/measurement.dto';
import { mapModelToRequestDto, mapResponseDtoToModel } from './mappers/measurement.mapper';

@Injectable({
    providedIn: 'root',
})
export class MeasurementService extends BaseApiRepository<Measurement, MeasurementResponseDto, MeasurementRequestDto> {
    protected readonly url = environment.backendGoUrl;

    protected readonly path = 'rooms/{roomId}/measurements';

    public mapResponseDtoToModel = mapResponseDtoToModel;

    public mapModelToRequestDto = mapModelToRequestDto;

    public getLatestByRoomId(roomId: number): Observable<ApiResponse<Measurement>> {
        let url = this.urlCombine();
        url = url.replace('{roomId}', roomId.toString());
        url = url + '/latest';

        return this.get(url);

        //        return this.get(url).pipe(
        //            map((apiResponse: ApiResponse<Temperature>) => {
        //                const temperature = apiResponse.data!;
        //                return temperature;
        //            }),
        //        );
    }
}
