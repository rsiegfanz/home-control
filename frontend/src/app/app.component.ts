import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule, MatIconRegistry } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { DomSanitizer } from '@angular/platform-browser';
import { RouterLink, RouterOutlet } from '@angular/router';
import { FlexModule } from '@ngbracket/ngx-layout/flex';

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [FlexModule, RouterLink, RouterOutlet, MatIconModule, MatToolbarModule, MatButtonModule],
    template: `<mat-toolbar color="primary">
            <button mat-icon-button><mat-icon>menu</mat-icon></button>
            <mat-icon svgIcon="home"></mat-icon>
            <a mat-button routerLink="/home"><h1>Home</h1></a>
            <span class="flex-spacer"></span>
            <button mat-mini-fab routerLink="/user/profile" matTooltip="Profile" aria-label="User Profile">
                <mat-icon>account_circle</mat-icon>
            </button>
            <button mat-mini-fab routerLink="/user/logout" matTooltip="Logout" aria-label="Logout">
                <mat-icon>lock_open</mat-icon>
            </button>
        </mat-toolbar>
        <router-outlet></router-outlet> `,
    styles: '',
})
export class AppComponent {
    constructor(iconRegistry: MatIconRegistry, sanitizer: DomSanitizer) {
        iconRegistry.addSvgIcon('home', sanitizer.bypassSecurityTrustResourceUrl('assets/img/icons/home.svg'));
    }
}

