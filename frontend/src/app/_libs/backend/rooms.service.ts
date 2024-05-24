import { Injectable } from '@angular/core';
import { BackendService } from './backend.service';
import Room from '../models/room.model';

@Injectable({
    providedIn: 'root',
})
export class RoomsService extends BackendService<Room> {
    protected override subdirectory = 'rooms';
}

