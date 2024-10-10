import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RoomId } from '../../_libs/house/enums/rooms.enum';
import Room from '../../_libs/house/models/room.model';
import { IconDataprovider } from '../../_libs/icons/icon.dataprovider';
import { NavBarRoomComponent } from './nav-bar-room/nav-bar-room.component';
import { RoomsViewService } from '../../_libs/house/viewServices/rooms.viewservice';

@Component({
    selector: 'app-nav-bar-rooms',
    standalone: true,
    imports: [CommonModule, NavBarRoomComponent],
    templateUrl: './nav-bar-rooms.component.html',
    styleUrl: './nav-bar-rooms.component.scss',
})
export class NavBarRoomsComponent {
    constructor(protected roomsService: RoomsViewService) {}
}
