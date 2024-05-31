import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RoomId } from '../../_libs/house/enums/rooms.enum';
import Room from '../../_libs/house/models/room.model';
import { IconDataprovider } from '../../_libs/icons/icon.dataprovider';
import { NavBarRoomComponent } from './nav-bar-room/nav-bar-room.component';

@Component({
    selector: 'app-nav-bar-rooms',
    standalone: true,
    imports: [CommonModule, NavBarRoomComponent],
    templateUrl: './nav-bar-rooms.component.html',
    styleUrl: './nav-bar-rooms.component.scss',
})
export class NavBarRoomsComponent {
    rooms: Room[] = [
        new Room(RoomId.GARAGE, 'outside', IconDataprovider.outside),
        new Room(RoomId.FIRST_FLOOR_GALLERY, 'Empore', IconDataprovider.gallery),
        new Room(RoomId.FIRST_FLOOR_LIVING_ROOM, 'Wohnzimmer', IconDataprovider.livingRoom),
        new Room(RoomId.FIRST_FLOOR_KITCHEN, 'Küche', IconDataprovider.kitchen),
        new Room(RoomId.FIRST_FLOOR_BEDROOM, 'Schlafzimmer', IconDataprovider.bedroom),
        new Room(RoomId.BASEMENT_GYM, 'Gym', IconDataprovider.gym),
    ];
}
