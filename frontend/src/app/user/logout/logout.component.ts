import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../auth/auth.service';

@Component({
    selector: 'app-logout',
    standalone: true,
    imports: [],
    template: ` <p>Logging out...</p> `,
    styles: ``,
})
export class LogoutComponent implements OnInit {
    constructor(
        private router: Router,
        private authService: AuthService,
    ) {}

    ngOnInit() {
        this.authService.logout(true);
        this.router.navigate(['/']);
    }
}

