import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { ApiResponse } from '../../backend/models/api-response.model';
import { BaseApiRepository } from '../../backend/services/repositories/base-api.repository';
import Temperature from '../models/temperature.model';
import { TemperatureRequestDto, TemperatureResponseDto } from './dtos/temperature.dto';
import { mapModelToRequestDto, mapResponseDtoToModel } from './mappers/temperature.mapper';

@Injectable({
    providedIn: 'root',
})
export class TemperatureService extends BaseApiRepository<Temperature, TemperatureResponseDto, TemperatureRequestDto> {
    protected readonly url = environment.backendGoUrl;

    protected readonly path = 'rooms/{roomId}/temperatures';

    public mapResponseDtoToModel = mapResponseDtoToModel;

    public mapModelToRequestDto = mapModelToRequestDto;

    public getLatestByRoomId(roomId: number): Observable<ApiResponse<Temperature>> {
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
