import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-page-not-found',
    standalone: true,
    imports: [RouterLink],
    template: `<p>This page does not exist. Back to <a routerLink="/home">Home</a></p>`,
    styles: ``,
})
export class PageNotFoundComponent {}

