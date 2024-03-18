import { HttpClient } from '@angular/common/http';
import urlJoin from 'url-join';

export abstract class BackendService {
    protected abstract backendUrl: string;

    constructor(protected http: HttpClient) {}

    protected createUrl(path: string): string {
        if (!path) {
            return this.backendUrl;
        }

        return urlJoin(this.backendUrl, path);
    }
}

