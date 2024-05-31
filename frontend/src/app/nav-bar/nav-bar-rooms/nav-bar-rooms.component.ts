import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import Room from '../../_libs/house/models/room.model';
import { NavBarRoomComponent } from './nav-bar-room/nav-bar-room.component';
import { RoomId } from '../../_libs/house/enums/rooms.enum';

@Component({
    selector: 'app-nav-bar-rooms',
    standalone: true,
    imports: [CommonModule, NavBarRoomComponent],
    templateUrl: './nav-bar-rooms.component.html',
    styleUrl: './nav-bar-rooms.component.scss',
})
export class NavBarRoomsComponent {
    rooms: Room[] = [
        { id: RoomId.GARAGE, name: 'outside' },
        { id: RoomId.FIRST_FLOOR_GALLERY, name: 'Empore' },
        { id: RoomId.FIRST_FLOOR_LIVING_ROOM, name: 'Wohnzimmer' },
        { id: RoomId.FIRST_FLOOR_KITCHEN, name: 'Küche' },
        { id: RoomId.FIRST_FLOOR_BEDROOM, name: 'Schlafzimmer' },
        { id: RoomId.BASEMENT_GYM, name: 'Gym' },
    ];
}
