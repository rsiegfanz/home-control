import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import Room from '../../_libs/house/models/room.model';
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
        { id: 1, name: 'outside' },
        { id: 2, name: 'bedroom' },
        { id: 3, name: 'kitchen' },
        { id: 4, name: 'office' },
    ];
}
