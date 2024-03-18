import { Role } from '../../auth/auth.enum';
import { IName } from './name.interface';
import { IPhone } from './phone.interface';

export interface IUser {
    _id: string;
    email: string;
    name: IName;
    picture: string;
    role: Role | string;
    userStatus: boolean;
    dateOfBirth: Date | null | string;
    level: number;
    address: {
        line1: string;
        line2?: string;
        city: string;
        state: string;
        zip: string;
    };
    phones: IPhone[];
    readonly fullName?: string;
}

