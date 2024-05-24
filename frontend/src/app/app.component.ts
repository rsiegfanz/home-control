import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { NavBarComponent } from './nav-bar/nav-bar.component';

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [NavBarComponent, RouterLink, RouterOutlet],
    templateUrl: './app.component.html',
    styles: '',
})
export class AppComponent {}
