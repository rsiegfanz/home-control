import { Component, Input } from '@angular/core';
import { TemperatureService } from '../../../_libs/house/backend/temperature.service';
import Room from '../../../_libs/house/models/room.model';

@Component({
    selector: 'app-dashboard-rooms',
    standalone: true,
    imports: [],
    templateUrl: './dashboard-rooms.component.html',
    styleUrl: './dashboard-rooms.component.scss',
})
export class DashboardRoomsComponent {
    @Input({ required: true }) rooms!: Room[];

    constructor(private _temperatureService: TemperatureService) {}

    ngOnInit() {
        // this._temperatureService.getLatestByRoomId(1).subscribe();
    }
}
