import { BaseModel } from '../../backend/models/base.model';

export default class Measurement extends BaseModel {
    public readonly recordedAt: Date;

    public readonly temperature: number;

    public readonly humidity: number;

    constructor(timestamp: Date, temperature: number, humidity: number) {
        super();
        this.recordedAt = timestamp;
        this.temperature = temperature;
        this.humidity = humidity;
    }
}
