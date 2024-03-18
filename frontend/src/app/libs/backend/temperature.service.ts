import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import Temperature from '../models/temperature.model';
import { BackendService } from './backend.service';

@Injectable({
    providedIn: 'root',
})
export class TemperatureService extends BackendService<Temperature> {
    protected override subdirectory = 'rooms/{roomId}/temperatures';

    public override getAll(): Observable<Temperature[]> {
        throw new Error('Method not implemented.');
    }

    public getLatestByRoomId(roomId: number): Observable<Temperature> {
        let url = this.createUrl();
        url = url.replace('{roomId}', roomId.toString());

        return this.get(url);
    }
}

