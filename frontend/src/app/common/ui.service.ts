import { importProvidersFrom, Injectable, makeEnvironmentProviders } from '@angular/core';
import { Observable } from 'rxjs';

import { SimpleDialogComponent } from './simple-dialog.component';

@Injectable({
    providedIn: 'root',
})
export class UiService {
    constructor() {}

    //    showToast(message: string, action = 'Close', config?: MatSnackBarConfig) {}

    //    showDialog(title: string, content: string, okText = 'OK', cancelText?: string, customConfig?: MatDialogConfig): Observable<boolean> {}
}
