import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RoomsService } from '../../_libs/house/backend/rooms.service';
import { DashboardRoomsComponent } from './dashboard-rooms/dashboard-rooms.component';

@Component({
    selector: 'app-dashboard-home',
    standalone: true,
    imports: [CommonModule, DashboardRoomsComponent],
    templateUrl: './dashboard-home.component.html',
    styleUrl: './dashboard-home.component.scss',
})
export class DashboardHomeComponent {
    constructor(private _roomsService: RoomsService) {}

    ngOnInit() {}
}
