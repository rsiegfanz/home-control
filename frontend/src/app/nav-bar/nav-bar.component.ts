import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ThemeSwitchComponent } from '../_libs/components/theme-switch/theme-switch.component';
import { NavBarRoomsComponent } from './nav-bar-rooms/nav-bar-rooms.component';

@Component({
    selector: 'app-nav-bar',
    standalone: true,
    imports: [CommonModule, NavBarRoomsComponent, ThemeSwitchComponent],
    templateUrl: './nav-bar.component.html',
    styleUrl: './nav-bar.component.scss',
})
export class NavBarComponent {
    ngOnInit(): void {}
}
