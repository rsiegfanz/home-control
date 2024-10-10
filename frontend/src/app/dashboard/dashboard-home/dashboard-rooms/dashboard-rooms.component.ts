import { Component, Input } from '@angular/core';
import { MeasurementService } from '../../../_libs/house/backend/services/measurement.service';
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

    constructor(private _measurementService: MeasurementService) {}

    ngOnInit() {}
}
