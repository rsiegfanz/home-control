import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { Router, RouterLink } from '@angular/router';
import { FlexModule } from '@ngbracket/ngx-layout';
import { AuthService } from '../auth/auth.service';
import { combineLatest, filter, tap } from 'rxjs';

@Component({
    selector: 'app-home',
    standalone: true,
    imports: [FlexModule, MatButtonModule, RouterLink],
    template: `
        <div fxLayout="column" fxLayoutAlign="center center">
            <span class="mat-headline-3">Hausübersicht</span>
            <button mat-raised-button color="primary" (click)="login()">Login</button>
            <button mat-raised-button color="primary" routerLink="/dashboard">Dashboard</button>
        </div>
    `,
    styles: `
        div[fxLayout] {
            margin-top: 32px;
        }
    `,
})
export class HomeComponent {
    constructor(
        private authService: AuthService,
        private router: Router,
    ) {}

    login(): void {
        this.authService.login('rs.88.tech@gmail.com', 'x');

        combineLatest([this.authService.authStatus$, this.authService.currentUser$])
            .pipe(
                filter(([authStatus, user]) => authStatus.isAuthenticated && user?._id !== ''),
                tap(([authStatus, user]) => {
                    this.router.navigate(['/dashboard']);
                }),
            )
            .subscribe();
    }
}

