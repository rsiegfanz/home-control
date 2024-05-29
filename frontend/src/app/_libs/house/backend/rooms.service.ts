import { Injectable } from '@angular/core';
import Room from '../models/room.model';
import { BackendService } from './backend.service';

@Injectable({
    providedIn: 'root',
})
export class RoomsService extends BackendService<Room> {
    protected override subdirectory = 'rooms';
}
