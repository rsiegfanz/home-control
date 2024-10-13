import { Component, OnInit } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { MeasurementService } from '../../_libs/house/backend/services/measurement.service';
import Measurement from '../../_libs/house/models/measurement.model';
import { ChartModule } from 'primeng/chart';

@Component({
    selector: 'app-climate-measurements-home',
    standalone: true,
    imports: [ChartModule],
    templateUrl: './climate-measurements-home.component.html',
    styleUrl: './climate-measurements-home.component.scss',
})
export class ClimateMeasurementsHomeComponent implements OnInit {
    chartData: any;

    chartOptions: any;

    climateData: Measurement[] | undefined;

    constructor(
        private apollo: Apollo,
        private _measurementService: MeasurementService,
    ) {}

    ngOnInit() {
        const startDate = '2024-10-11T00:00:00Z';
        const endDate = '2024-10-11T23:59:59Z';
        const roomExternalId = 'e868e758ea4f';

        this.query(startDate, endDate, roomExternalId);

        // this.updateGraph();
    }

    public query(startDate: string, endDate: string, roomExternalId: string): void {
        this._measurementService.query(startDate, endDate, roomExternalId).subscribe((result) => {
            this.climateData = result;
            this.updateGraph(result);
        });
    }

    public updateGraph(climateMeasurements: Measurement[]): void {
        const temperatures = climateMeasurements.map((measurement: any) => measurement.temperature);
        const humidity = climateMeasurements.map((measurement: any) => measurement.humidity);
        const recordedAt = climateMeasurements.map((measurement: any) => new Date(measurement.recordedAt).toLocaleDateString());

        this.chartData = {
            labels: recordedAt,
            datasets: [
                {
                    label: 'Temperatur (°C)',
                    data: temperatures,
                    fill: false,
                    borderColor: '#42A5F5',
                    tension: 0.4,
                },
                {
                    label: 'Luftfeuchtigkeit (%)',
                    data: humidity,
                    fill: false,
                    borderColor: '#66BB6A',
                    tension: 0.4,
                },
            ],
        };

        this.chartOptions = {
            responsive: true,
            plugins: {
                legend: {
                    display: true,
                    position: 'top',
                },
            },
            // scales: {
            //     x: {
            //         type: 'time',
            //         time: {
            //             unit: 'minute', // Beispiel für Minutengenauigkeit
            //         },
            //     },
            // },
        };
    }

    // public updateGraph2(): void {
    //     const documentStyle = getComputedStyle(document.documentElement);
    //     const textColor = documentStyle.getPropertyValue('--text-color');
    //     const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    //     const surfaceBorder = documentStyle.getPropertyValue('--surface-border');

    //     this.chartData = {
    //         labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    //         datasets: [
    //             {
    //                 label: 'First Dataset',
    //                 data: [65, 59, 80, 81, 56, 55, 40],
    //                 fill: false,
    //                 borderColor: documentStyle.getPropertyValue('--blue-500'),
    //                 tension: 0.4,
    //             },
    //             {
    //                 label: 'Second Dataset',
    //                 data: [28, 48, 40, 19, 86, 27, 90],
    //                 fill: false,
    //                 borderColor: documentStyle.getPropertyValue('--pink-500'),
    //                 tension: 0.4,
    //             },
    //         ],
    //     };

    //     this.chartOptions = {
    //         maintainAspectRatio: false,
    //         aspectRatio: 0.6,
    //         plugins: {
    //             legend: {
    //                 labels: {
    //                     color: textColor,
    //                 },
    //             },
    //         },
    //         scales: {
    //             x: {
    //                 ticks: {
    //                     color: textColorSecondary,
    //                 },
    //                 grid: {
    //                     color: surfaceBorder,
    //                     drawBorder: false,
    //                 },
    //             },
    //             y: {
    //                 ticks: {
    //                     color: textColorSecondary,
    //                 },
    //                 grid: {
    //                     color: surfaceBorder,
    //                     drawBorder: false,
    //                 },
    //             },
    //         },
    //     };
    // }
}
