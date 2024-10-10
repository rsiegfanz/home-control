import { BaseModel } from '../../backend/models/base.model';

export default class Measurement extends BaseModel {
    public readonly timestamp: Date;

    public readonly temperature: number;

    public readonly humidity: number;

    constructor(timestamp: Date, temperature: number, humidity: number) {
        super();
        this.timestamp = timestamp;
        this.temperature = temperature;
        this.humidity = humidity;
    }
}
