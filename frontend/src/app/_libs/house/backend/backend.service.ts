import { HttpClient } from '@angular/common/http';
import { inject } from '@angular/core';
import { Observable, map } from 'rxjs';
import urlJoin from 'url-join';
import { environment } from '../../../environments/environment';

export abstract class BackendService<T> {
    protected http = inject(HttpClient);

    protected baseUrl = environment.backendGoUrl;

    protected abstract subdirectory: string;

    public get(url?: string): Observable<T> {
        const queryUrl = url ?? this.createUrl();

        return this.http.get<T>(queryUrl).pipe(
            map((data) => {
                console.log(data);
                return data;
            }),
        );
    }

    public getAll(url?: string): Observable<T[]> {
        const queryUrl = url ?? this.createUrl();

        return this.http.get<T[]>(queryUrl).pipe(
            map((data) => {
                console.log(data);
                return data;
            }),
        );
    }

    protected createUrl(path?: string): string {
        if (!path) {
            return urlJoin(this.baseUrl, this.subdirectory);
        }

        return urlJoin(this.baseUrl, this.subdirectory, path);
    }
}
