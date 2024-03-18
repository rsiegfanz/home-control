import { Role } from '../../auth/auth.enum';
import { IName } from './name.interface';
import { IPhone } from './phone.interface';
import { IUser } from './user.interface';
import { $enum } from 'ts-enum-util';

export class User implements IUser {
    constructor(
        // tslint:disable-next-line: variable-name
        public _id = '',
        public email = '',
        public name = { first: '', middle: '', last: '' } as IName,
        public picture = '',
        public role = Role.None,
        public dateOfBirth: Date | null = null,
        public userStatus = false,
        public level = 0,
        public address = {
            line1: '',
            city: '',
            state: '',
            zip: '',
        },
        public phones: IPhone[] = [],
    ) {}

    static Build(user: IUser) {
        if (!user) {
            return new User();
        }

        return new User(
            user._id,
            user.email,
            user.name,
            user.picture,
            $enum(Role).asValueOrDefault(user.role, Role.None),
            typeof user.dateOfBirth === 'string' ? new Date(user.dateOfBirth) : user.dateOfBirth,
            user.userStatus,
            user.level,
            user.address,
            user.phones,
        );
    }

    public get fullName(): string {
        if (!this.name) {
            return '';
        }

        if (this.name.middle) {
            return `${this.name.first} ${this.name.middle} ${this.name.last}`;
        }
        return `${this.name.first} ${this.name.last}`;
    }

    toJSON(): object {
        const serialized = Object.assign(this);
        delete serialized._id;
        delete serialized.fullName;
        return serialized;
    }
}

