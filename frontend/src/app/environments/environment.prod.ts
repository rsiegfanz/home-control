import { AuthMode } from '../auth/auth-mode.enum';

export const environment = {
    production: false,
    authMode: AuthMode.InMemory,
    baseUrl: '',
    firebase: {
        apiKey: '',
        authDomain: '',
        databaseURL: '',
        projectId: '',
        storageBucket: '',
        messagingSenderId: '',
        appId: '',
    },
};

