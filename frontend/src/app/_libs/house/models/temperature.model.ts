export default class Temperature {
    public readonly timestamp: Date;

    public readonly value: number;

    constructor(timestamp: Date, value: number) {
        this.timestamp = timestamp;
        this.value = value;
    }
}
