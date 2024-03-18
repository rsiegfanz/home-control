import { BehaviorSubject, Observable } from 'rxjs';
import { IUser } from '../user/user/user.interface';
import { IAuthStatus } from './auth-status.interface';

export interface IAuthService {
    readonly authStatus$: BehaviorSubject<IAuthStatus>;
    readonly currentUser$: BehaviorSubject<IUser>;
    login(email: string, password: string): Observable<void>;
    logout(clearToken?: boolean): void;
    getToken(): string;
}

