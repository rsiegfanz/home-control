import { PhoneType } from './phone-type.enum';

export interface IPhone {
    type: PhoneType;
    digits: string;
    id: number;
}

