import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class ConfigService {
    private readonly _ip = environment.backendAiUrl;

    public readonly backendUrlStt = `${this._ip}/api/v2/stt`;

    public readonly backendUrlTts = `${this._ip}/api/v2/tts`;

    public readonly backendUrlChat = `${this._ip}/api/v2/ai/message`;

    public readonly backendUrlPersistence = `${this._ip}/api/v2/persistence/`;
}
