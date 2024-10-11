export interface IMeasurementDto {
    timestamp: Date;
    temperature: number;
    humidity: number;
}

export interface IClimateMeasurementSchemaRootDto {
    climateMeasurements: IClimateMeasurementSchemaDto[];
}

export interface IClimateMeasurementSchemaDto {
    __typename: string;
    roomExternalId: string;
    recordedAt: string;
    temperature: number;
    humidity: number;
}
