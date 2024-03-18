import { inject } from '@angular/core';
import { jwtDecode as decode } from 'jwt-decode';
import { BehaviorSubject, Observable, catchError, filter, map, mergeMap, pipe, tap, throwError } from 'rxjs';
import { CacheService } from '../common/cache.service';
import { transformError } from '../common/common';
import { IUser } from '../user/user/user.interface';
import { User } from '../user/user/user.model';
import { IAuthService } from './auth-service.interface';
import { IAuthStatus } from './auth-status.interface';
import { defaultAuthStatus } from './default-auth-status.const';
import { IServerAuthResponse } from './server-auth-response.interface';

export abstract class AuthService implements IAuthService {
    private getAndUpdateUserIfAuthenticated = pipe(
        filter((status: IAuthStatus) => status.isAuthenticated),
        mergeMap(() => this.getCurrentUser()),
        map((user: IUser) => this.currentUser$.next(user)),
        catchError(transformError),
    );

    protected readonly cache = inject(CacheService);

    readonly authStatus$ = new BehaviorSubject<IAuthStatus>(defaultAuthStatus);

    readonly currentUser$ = new BehaviorSubject<IUser>(new User());

    protected readonly resumeCurrentUser$ = this.authStatus$.pipe(this.getAndUpdateUserIfAuthenticated);

    // Example caching technique
    // readonly authStatus$ = new BehaviorSubject<IAuthStatus>(
    //   this.getItem('authStatus') ?? defaultAuthStatus
    // )

    constructor() {
        // this.authStatus$.pipe(tap((authStatus) => this.setItem('authStatus', authStatus))) // example caching technique

        if (this.hasExpiredToken()) {
            this.logout(true);
        } else {
            this.authStatus$.next(this.getAuthStatusFromToken());
            // To load user on browser refresh, resume pipeline must activate on the next cycle
            // Which allows for all services to constructed properly
            setTimeout(() => this.resumeCurrentUser$.subscribe(), 0);
        }
    }

    protected abstract authProvider(email: string, password: string): Observable<IServerAuthResponse>;
    protected abstract transformJwtToken(token: unknown): IAuthStatus;
    protected abstract getCurrentUser(): Observable<User>;

    login(email: string, password: string): Observable<void> {
        this.clearToken();

        const loginResponse$ = this.authProvider(email, password).pipe(
            map((value) => {
                console.log('SET TOKEN');
                this.setToken(value.accessToken);
                // const token = decode(value.accessToken)
                // return this.transformJwtToken(token)
                return this.getAuthStatusFromToken(); // Keeping the code DRY!
            }),
            tap((status) => this.authStatus$.next(status)),
            // filter((status: IAuthStatus) => status.isAuthenticated),
            // mergeMap(() => this.getCurrentUser()),
            // map((user: IUser) => this.currentUser$.next(user)),
            // catchError(transformError)
            this.getAndUpdateUserIfAuthenticated, // Keeping the code DRY!
        );

        loginResponse$.subscribe({
            error: (err) => {
                console.log('ERR', err);
                this.logout();
                return throwError(() => err);
            },
        });

        return loginResponse$;
    }

    logout(clearToken?: boolean) {
        if (clearToken) {
            this.clearToken();
        }
        setTimeout(() => this.authStatus$.next(defaultAuthStatus), 0);
    }

    protected setToken(jwt: string) {
        this.cache.setItem('jwt', jwt);
    }

    getToken(): string {
        return this.cache.getItem('jwt') ?? '';
    }

    protected clearToken() {
        // clear all cache data along with jwt
        this.cache.clear();
    }

    protected hasExpiredToken(): boolean {
        const jwt = this.getToken();

        if (jwt) {
            const payload = decode(jwt);
            return Date.now() >= payload.exp! * 1000;
        }

        return true;
    }

    protected getAuthStatusFromToken(): IAuthStatus {
        return this.transformJwtToken(decode(this.getToken()));
    }
}

