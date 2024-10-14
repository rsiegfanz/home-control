import { Component, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { MeasurementService } from '../../_libs/house/backend/services/measurement.service';
import Measurement from '../../_libs/house/models/measurement.model';
import { ChartModule } from 'primeng/chart';
import { DropdownModule } from 'primeng/dropdown';
import 'chartjs-adapter-luxon';
import { RoomsViewService } from '../../_libs/house/viewServices/rooms.viewservice';
import Room from '../../_libs/house/models/room.model';
import { LoadingDataSpinnerComponent } from '../../_libs/components/loading-data-spinner/loading-data-spinner.component';
import { CommonModule } from '@angular/common';
import { LoadingStatus } from '../../_libs/house/enums/loading-status.enum';
import { CalendarModule } from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import { Subscription } from 'rxjs';

@Component({
    selector: 'app-climate-measurements-home',
    standalone: true,
    imports: [CommonModule, CalendarModule, ChartModule, DropdownModule, LoadingDataSpinnerComponent, FormsModule],
    templateUrl: './climate-measurements-home.component.html',
    styleUrl: './climate-measurements-home.component.scss',
})
export class ClimateMeasurementsHomeComponent implements OnInit, OnDestroy {
    public chartData: any;

    public chartOptions: any;

    public climateData: Measurement[] | undefined;

    private measurementSubscription: Subscription | null = null;

    public startDate: Date;
    public endDate: Date;

    public selectedRoom: Room | undefined;

    public readonly loadingRoomsText = 'Räume werden geladen';
    public errorRoomsText = '';

    public readonly loadingClimateMeasurementsText = 'Daten werden geladen';
    public loadingClimateMeasurementsError = '';

    public loadingStatusClimateMeasurements: WritableSignal<LoadingStatus> = signal(LoadingStatus.NONE);

    constructor(
        public roomViewService: RoomsViewService,
        private _measurementService: MeasurementService,
    ) {
        const today = new Date();
        this.startDate = new Date(today.getFullYear(), today.getMonth(), today.getDate());
        this.endDate = new Date(today.getFullYear(), today.getMonth(), today.getDate(), 23, 59, 59);
    }

    ngOnInit() {}

    ngOnDestroy() {
        if (this.measurementSubscription) {
            this.measurementSubscription.unsubscribe();
        }
    }

    public onRoomChange(event: { value: Room }): void {
        if (!event || !event.value) {
            return;
        }

        this.selectedRoom = event.value;
        this._query(this.startDate, this.endDate, this.selectedRoom!.externalRoomId);
        this._subscribeToMeasurements(this.selectedRoom.externalRoomId);
    }

    public onDateChange(): void {
        console.log('datechange');
        if (this.startDate > this.endDate) {
            const temp = this.startDate;
            this.startDate = this.endDate;
            this.endDate = temp;
        }

        console.log(this.startDate);
        console.log(this.endDate);

        if (this.selectedRoom) {
            console.log('zzz');
            this._query(this.startDate, this.endDate, this.selectedRoom.externalRoomId);
        }
    }

    private _query(startDate: Date, endDate: Date, roomExternalId: string): void {
        this.loadingStatusClimateMeasurements.set(LoadingStatus.LOADING);
        this.loadingClimateMeasurementsError = '';

        this._measurementService.query(startDate.toISOString(), endDate.toISOString(), roomExternalId).subscribe({
            next: (result) => {
                this.climateData = result;
                this.updateGraph(result);
                if (result.length <= 0) {
                    this.loadingStatusClimateMeasurements.set(LoadingStatus.LOADING_ERROR);
                    this.loadingClimateMeasurementsError = 'Keine Daten vorhanden';
                } else {
                    this.loadingStatusClimateMeasurements.set(LoadingStatus.LOADING_SUCCESS);
                }
            },
            error: (error) => {
                // todo: error handling
                console.error('Error fetching climate data:', error);
                this.loadingStatusClimateMeasurements.set(LoadingStatus.LOADING_ERROR);
                this.loadingClimateMeasurementsError = 'Fehler beim Laden der Daten';
            },
        });
    }

    private _subscribeToMeasurements(roomExternalId: string): void {
        if (this.measurementSubscription) {
            this.measurementSubscription.unsubscribe();
        }

        this.measurementSubscription = this._measurementService.subscribeToMeasurements(roomExternalId).subscribe({
            next: (newMeasurement) => {
                console.log('New measurement received:', newMeasurement);
                // Hier können Sie die Logik implementieren, um die neuen Daten
                // in this.climateData einzufügen und updateGraph aufzurufen
                if (this.climateData) {
                    this.climateData.push(newMeasurement);
                    this.updateGraph(this.climateData);
                }
            },
            error: (error) => {
                console.error('Error in climate measurement subscription:', error);
            },
        });
    }

    public updateGraph(climateMeasurements: Measurement[]): void {
        if (climateMeasurements?.length < 0) {
            // todo error handling
            return;
        }

        const temperatures = climateMeasurements.map((measurement) => ({
            x: measurement.recordedAt,
            y: measurement.temperature,
        }));
        const humidity = climateMeasurements.map((measurement) => ({
            x: measurement.recordedAt,
            y: measurement.humidity,
        }));
        const recordedAt = climateMeasurements.map((measurement) => measurement.recordedAt);

        const minDate = new Date(Math.min(...recordedAt.map((d) => d.getTime())));
        const maxDate = new Date(Math.max(...recordedAt.map((d) => d.getTime())));

        this.chartData = {
            datasets: [
                {
                    label: 'Temperatur (°C)',
                    data: temperatures,
                    fill: false,
                    borderColor: '#42A5F5',
                    tension: 0.4,
                    yAxisID: 'y-axis-temperature',
                },
                {
                    label: 'Luftfeuchtigkeit (%)',
                    data: humidity,
                    fill: false,
                    borderColor: '#66BB6A',
                    tension: 0.4,
                    yAxisID: 'y-axis-humidity',
                },
            ],
        };

        this.chartOptions = {
            responsive: true,
            interaction: {
                mode: 'index',
                intersect: false,
            },
            plugins: {
                legend: {
                    display: true,
                    position: 'top',
                },
            },
            scales: {
                x: {
                    type: 'time',
                    time: {
                        unit: 'hour',
                        displayFormats: {
                            hour: 'HH:mm',
                        },
                        tooltipFormat: 'dd.MM.yyyy HH:mm',
                    },
                    title: {
                        display: true,
                        text: 'Uhrzeit',
                    },
                    adapters: {
                        date: {
                            locale: 'de-DE',
                        },
                    },
                    min: minDate,
                    max: maxDate,
                },
                'y-axis-temperature': {
                    type: 'linear',
                    display: true,
                    position: 'left',
                    title: {
                        display: true,
                        text: 'Temperatur (°C)',
                    },
                    grid: {
                        drawOnChartArea: true,
                    },
                },
                'y-axis-humidity': {
                    type: 'linear',
                    display: true,
                    position: 'right',
                    title: {
                        display: true,
                        text: 'Luftfeuchtigkeit (%)',
                    },
                    grid: {
                        drawOnChartArea: false,
                    },
                },
            },
        };
    }
}
