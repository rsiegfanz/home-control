import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { FaIconComponent } from '@fortawesome/angular-fontawesome';
import { Subscription, interval, startWith, switchMap } from 'rxjs';
import { ApiResponse } from '../../../_libs/backend/models/api-response.model';
import { MeasurementService } from '../../../_libs/house/backend/measurement.service';
import Measurement from '../../../_libs/house/models/measurement.model';
import Room from '../../../_libs/house/models/room.model';
import { IconDataprovider } from '../../../_libs/icons/icon.dataprovider';

@Component({
    selector: 'app-nav-bar-room',
    standalone: true,
    imports: [CommonModule, FaIconComponent],
    templateUrl: './nav-bar-room.component.html',
    styleUrl: './nav-bar-room.component.scss',
})
export class NavBarRoomComponent {
    @Input() room!: Room;

    measurement: Measurement | undefined;

    public iconProvider = IconDataprovider;

    private _timeInterval: Subscription | undefined;

    private readonly INTERVAL = 3000;

    constructor(private _measurementService: MeasurementService) {}

    ngOnInit(): void {
        this.getTemperature();
    }

    getTemperature() {
        this._timeInterval = interval(this.INTERVAL)
            .pipe(
                startWith(0),
                switchMap(() => this._measurementService.getLatestByRoomId(this.room.id)),
            )
            .subscribe((apiResponse: ApiResponse<Measurement>) => {
                if (apiResponse.isError) {
                    return;
                }

                this.measurement = apiResponse.data!;
            });
    }

    ngOnDestroy(): void {
        this._timeInterval?.unsubscribe();
    }
}
