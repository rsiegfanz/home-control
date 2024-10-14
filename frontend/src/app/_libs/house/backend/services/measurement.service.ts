import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ApiResponse } from '../../../backend/models/api-response.model';
import Measurement from '../../models/measurement.model';
import { mapDtoToModel, mapGraphlQLDtoToModel } from '../mappers/measurement.mapper';
import { BackendService } from '../../../backend/services/backend.service';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { IClimateMeasurementSchemaDto, IClimateMeasurementSchemaRootDto } from '../dtos/measurement.dto';
import { IApiGraphQLResponseDto } from '../../../backend/dtos/api-graphql-response-dto.interface';

const GET_CLIMATE_MEASUREMENTS = gql`
    query GetClimateMeasurements($startDate: String!, $endDate: String!, $roomExternalId: String!) {
        climateMeasurements(startDate: $startDate, endDate: $endDate, roomExternalId: $roomExternalId) {
            recordedAt
            roomExternalId
            temperature
            humidity
        }
    }
`;

const CLIMATE_MEASUREMENT_SUBSCRIPTION = gql`
    subscription ClimateMeasurementUpdates($roomExternalId: String!) {
        climateMeasurementUpdates(roomExternalId: $roomExternalId) {
            recordedAt
            roomExternalId
            temperature
            humidity
        }
    }
`;

@Injectable({
    providedIn: 'root',
})
export class MeasurementService extends BackendService<Measurement> {
    protected override subdirectory = 'rooms/{roomId}/measurements';

    constructor(protected apollo: Apollo) {
        super();
    }

    public getLatestByExternalRoomId(externalRoomId: string): Observable<ApiResponse<Measurement>> {
        const url = this.createUrl().replace('{roomId}', externalRoomId) + '/latest';
        return this.getSingle(url, mapDtoToModel);
    }

    public query(startDate: string, endDate: string, roomExternalId: string): Observable<Measurement[]> {
        return this.apollo
            .watchQuery<IClimateMeasurementSchemaRootDto>({
                query: GET_CLIMATE_MEASUREMENTS,
                variables: { startDate, endDate, roomExternalId },
            })
            .valueChanges.pipe(
                map((result) => {
                    if (!result.data) {
                        throw new Error('No data returned from GraphQL query');
                    }
                    return result.data.climateMeasurements.map(mapGraphlQLDtoToModel);
                }),
            );
    }

    public subscribeToMeasurements(roomExternalId: string): Observable<Measurement> {
        return this.apollo
            .subscribe<{ climateMeasurementUpdates: IClimateMeasurementSchemaDto }>({
                query: CLIMATE_MEASUREMENT_SUBSCRIPTION,
                variables: { roomExternalId },
            })
            .pipe(
                map((result) => {
                    if (!result.data) {
                        throw new Error('No data returned from GraphQL subscription');
                    }
                    return mapGraphlQLDtoToModel(result.data.climateMeasurementUpdates);
                }),
            );
    }
}
