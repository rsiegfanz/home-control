import { Component, OnInit } from '@angular/core';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';

@Component({
    selector: 'app-climate-measurements-home',
    standalone: true,
    imports: [],
    templateUrl: './climate-measurements-home.component.html',
    styleUrl: './climate-measurements-home.component.scss',
})
export class ClimateMeasurementsHomeComponent implements OnInit {
    climateData: any[] | undefined;

    constructor(private apollo: Apollo) {}

    ngOnInit() {
        this.apollo
            .watchQuery({
                query: gql`
                    query GetClimateMeasurements($startDate: String!, $endDate: String!, $roomExternalId: String!) {
                        climateMeasurements(startDate: $startDate, endDate: $endDate, roomExternalId: $roomExternalId) {
                            recordedAt
                            roomExternalId
                            temperature
                            humidity
                        }
                    }
                `,
                variables: {
                    startDate: '2024-10-11T00:00:00Z',
                    endDate: '2024-10-11T23:59:59Z',
                    roomExternalId: 'e868e758ea4f',
                },
            })
            .valueChanges.subscribe((result: any) => {
                this.climateData = result?.data?.climateMeasurements;
            });
    }
}
