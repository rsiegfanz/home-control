import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ThemeSwitchComponent } from '../_libs/components/theme-switch/theme-switch.component';
import Room from '../_libs/models/room.model';
import { NavBarRoomComponent } from './nav-bar-room/nav-bar-room.component';

@Component({
    selector: 'app-nav-bar',
    standalone: true,
    imports: [CommonModule, NavBarRoomComponent, ThemeSwitchComponent],
    templateUrl: './nav-bar.component.html',
    styleUrl: './nav-bar.component.scss',
})
export class NavBarComponent {
    rooms: Room[] = [
        { id: 1, name: 'outside' },
        { id: 2, name: 'bedroom' },
        { id: 3, name: 'kitchen' },
        { id: 4, name: 'office' },
    ];

    ngOnInit(): void {}
}
