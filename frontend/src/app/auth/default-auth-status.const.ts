import { IAuthStatus } from './auth-status.interface';
import { Role } from './auth.enum';

export const defaultAuthStatus: IAuthStatus = {
    isAuthenticated: false,
    userRole: Role.None,
    userId: '',
};

