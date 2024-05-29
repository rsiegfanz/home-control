import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { RoomsService } from '../../_libs/house/backend/rooms.service';
import Room from '../../_libs/house/models/room.model';
import { DashboardRoomsComponent } from './dashboard-rooms/dashboard-rooms.component';

@Component({
    selector: 'app-dashboard-home',
    standalone: true,
    imports: [CommonModule, DashboardRoomsComponent],
    templateUrl: './dashboard-home.component.html',
    styleUrl: './dashboard-home.component.scss',
})
export class DashboardHomeComponent {
    public rooms$: Observable<Room[]>;

    constructor(private _roomsService: RoomsService) {
        this.rooms$ = this._roomsService.getAll();
    }

    ngOnInit() {}
}
