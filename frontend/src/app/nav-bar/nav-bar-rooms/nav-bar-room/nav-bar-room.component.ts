import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { Subscription, interval, startWith, switchMap } from 'rxjs';
import { ApiResponse } from '../../../_libs/backend/models/api-response.model';
import { TemperatureService } from '../../../_libs/house/backend/temperature.service';
import Room from '../../../_libs/house/models/room.model';
import Temperature from '../../../_libs/house/models/temperature.model';

@Component({
    selector: 'app-nav-bar-room',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './nav-bar-room.component.html',
    styleUrl: './nav-bar-room.component.scss',
})
export class NavBarRoomComponent {
    @Input() room!: Room;

    temperature: number | undefined;

    private _timeInterval: Subscription | undefined;

    private readonly INTERVAL = 3000;

    constructor(private _temperatureService: TemperatureService) {}

    ngOnInit(): void {
        this.getTemperature();
    }

    getTemperature() {
        this._timeInterval = interval(this.INTERVAL)
            .pipe(
                startWith(0),
                switchMap(() => this._temperatureService.getLatestByRoomId(this.room.id)),
            )
            .subscribe((apiResponse: ApiResponse<Temperature>) => {
                console.log('1');
                if (apiResponse.isError) {
                    console.log('2');
                    return;
                }

                this.temperature = apiResponse.data!.value;

                //    const val = Number(value.value);
                //  this.temperature = Number((Math.round(val * 100) / 100).toFixed(2));
            });
    }

    ngOnDestroy(): void {
        this._timeInterval?.unsubscribe();
    }
}
