import { Component, OnInit } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { MeasurementService } from '../../_libs/house/backend/services/measurement.service';
import Measurement from '../../_libs/house/models/measurement.model';

@Component({
    selector: 'app-climate-measurements-home',
    standalone: true,
    imports: [],
    templateUrl: './climate-measurements-home.component.html',
    styleUrl: './climate-measurements-home.component.scss',
})
export class ClimateMeasurementsHomeComponent implements OnInit {
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
    }

    query(startDate: string, endDate: string, roomExternalId: string) {
        this._measurementService.query(startDate, endDate, roomExternalId).subscribe((result) => {
            this.climateData = result;
        });
    }
}
