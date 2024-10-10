import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiResponse } from '../../../backend/models/api-response.model';
import Measurement from '../../models/measurement.model';
import { mapDtoToModel } from '../mappers/measurement.mapper';
import { BackendService } from '../../../backend/services/backend.service';

@Injectable({
    providedIn: 'root',
})
export class MeasurementService extends BackendService<Measurement> {
    protected override subdirectory = 'rooms/{roomId}/measurements';

    public getLatestByExternalRoomId(externalRoomId: string): Observable<ApiResponse<Measurement>> {
        let url = this.createUrl();
        url = url.replace('{roomId}', externalRoomId);
        url = url + '/latest';

        return this.getSingle(url, mapDtoToModel);
    }
}
