import { Role } from './auth.enum';

export interface IAuthStatus {
    isAuthenticated: boolean;
    userRole: Role;
    userId: string;
}

