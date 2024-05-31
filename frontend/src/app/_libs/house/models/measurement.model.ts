export default class Measurement {
    public readonly timestamp: Date;

    public readonly temperature: number;

    public readonly humidity: number;

    constructor(timestamp: Date, temperature: number, humidity: number) {
        this.timestamp = timestamp;
        this.temperature = temperature;
        this.humidity = humidity;
    }
}
