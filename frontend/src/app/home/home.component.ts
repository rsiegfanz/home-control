import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';

@Component({
    selector: 'app-home',
    standalone: true,
    imports: [RouterLink],
    template: `
        <div fxLayout="column" fxLayoutAlign="center center">
            <span class="">Hausübersicht</span>
            <button>Login</button>
            <button routerLink="/dashboard">Dashboard</button>
        </div>
    `,
    styles: `
        div {
            margin-top: 32px;
        }
    `,
})
export class HomeComponent {
    constructor(private router: Router) {}
}
