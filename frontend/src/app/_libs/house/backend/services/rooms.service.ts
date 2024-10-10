import { Injectable } from '@angular/core';
import Room from '../../models/room.model';
import { ApiResponse } from '../../../backend/models/api-response.model';
import { Observable } from 'rxjs';
import { mapDtoToModel, mapDtoToModelArray } from '../mappers/room.mapper';
import { BackendService } from '../../../backend/services/backend.service';

@Injectable({
    providedIn: 'root',
})
export class RoomsService extends BackendService<Room> {
    protected override subdirectory = 'rooms';

    public all(): Observable<ApiResponse<Room[]>> {
        const url = this.createUrl();

        return this.getMultiple(url, mapDtoToModelArray);
    }
}
