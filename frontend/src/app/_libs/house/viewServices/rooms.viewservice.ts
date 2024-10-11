import { Injectable } from '@angular/core';
import Room from '../models/room.model';
import { RoomsService } from '../backend/services/rooms.service';
import { ApiResponse } from '../../backend/models/api-response.model';

@Injectable({
    providedIn: 'root',
})
export class RoomsViewService {
    rooms: Room[] | undefined;

    constructor(private _roomsService: RoomsService) {
        this.rooms = undefined;

        // this._roomsService.all().subscribe((apiResonse: ApiResponse<Room[]>) => {
        //     if (apiResonse.isSuccessful) {
        //         this.rooms = apiResonse.data;
        //     }
        // });
    }
}
