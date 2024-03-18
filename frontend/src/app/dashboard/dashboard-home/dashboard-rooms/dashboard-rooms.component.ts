import { Component, Input } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { TemperatureService } from '../../../libs/backend/temperature.service';
import Room from '../../../libs/models/room.model';

@Component({
    selector: 'app-dashboard-rooms',
    standalone: true,
    imports: [MatCardModule],
    templateUrl: './dashboard-rooms.component.html',
    styleUrl: './dashboard-rooms.component.scss',
})
export class DashboardRoomsComponent {
    @Input({ required: true }) rooms!: Room[];

    constructor(private _temperatureService: TemperatureService) {}

    ngOnInit() {
        // this._temperatureService.getLatestByRoomId(roomId)
    }
}

