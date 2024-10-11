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

@Injectable({
    providedIn: 'root',
})
export class MeasurementService extends BackendService<Measurement> {
    protected override subdirectory = 'rooms/{roomId}/measurements';

    public constructor(protected apollo: Apollo) {
        super();
    }

    public getLatestByExternalRoomId(externalRoomId: string): Observable<ApiResponse<Measurement>> {
        let url = this.createUrl();
        url = url.replace('{roomId}', externalRoomId);
        url = url + '/latest';

        return this.getSingle(url, mapDtoToModel);
    }

    public query(startDate: string, endDate: string, roomExternalId: string): Observable<Measurement[]> {
        return this.apollo
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
                    startDate: startDate,
                    endDate: endDate,
                    roomExternalId: roomExternalId,
                },
            })
            .valueChanges.pipe(
                map((result: IApiGraphQLResponseDto) => {
                    if (!result || !result.data) {
                        console.log(`GraphQL Error: ${result}`);
                    }

                    const dtos = (result.data as IClimateMeasurementSchemaRootDto).climateMeasurements;
                    return dtos.map(mapGraphlQLDtoToModel);
                }),
            );
    }
}
