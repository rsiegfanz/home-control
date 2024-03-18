import { Injectable } from '@angular/core';
import { sign } from 'fake-jwt-sign';
import { Observable, of } from 'rxjs';
import { PhoneType } from '../user/user/phone-type.enum';
import { User } from '../user/user/user.model';
import { IAuthStatus } from './auth-status.interface';
import { Role } from './auth.enum';
import { AuthService } from './auth.service';
import { IServerAuthResponse } from './server-auth-response.interface';
@Injectable({
    providedIn: 'root',
})
export class InMemoryAuthService extends AuthService {
    private defaultUser = User.Build({
        _id: '5da01751da27cc462d265913',
        email: 'rs.88.tech@gmail.com',
        name: { first: 'Robert', last: 'S' },
        picture: 'https://secure.gravatar.com/avatar/7cbaa9afb5ca78d97f3c689f8ce6c985',
        role: Role.User,
        dateOfBirth: new Date(1980, 1, 1),
        userStatus: true,
        address: {
            line1: '101 Sesame St.',
            city: 'Bethesda',
            state: 'Maryland',
            zip: '20810',
        },
        level: 2,
        phones: [
            {
                id: 0,
                type: PhoneType.Mobile,
                digits: '5555550717',
            },
        ],
    });

    constructor() {
        super();
        console.warn("You're using the InMemoryAuthService. Do not use this service in production.");
    }

    protected authProvider(email: string, _password: string): Observable<IServerAuthResponse> {
        const authStatus = {
            isAuthenticated: true,
            userId: this.defaultUser._id,
            userRole: Role.User,
        } as IAuthStatus;

        this.defaultUser.role = authStatus.userRole;

        const authResponse = {
            accessToken: sign(authStatus, 'secret', {
                expiresIn: '1h',
                algorithm: 'none',
            }),
        } as IServerAuthResponse;

        return of(authResponse);
    }

    protected transformJwtToken(token: IAuthStatus): IAuthStatus {
        return token;
    }

    protected getCurrentUser(): Observable<User> {
        return of(this.defaultUser);
    }
}

