import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
    selector: 'app-theme-switch',
    standalone: true,
    imports: [CommonModule, FormsModule],
    templateUrl: './theme-switch.component.html',
    styleUrl: './theme-switch.component.scss',
})
export class ThemeSwitchComponent {
    themes = ['light', 'futuristic'];

    selectedTheme: 'light' | 'futuristic' = 'light';

    public selectChange($event: unknown) {
        if (!$event) {
            return;
        }
        console.log($event);
        document.querySelector('html')!.setAttribute('data-theme', $event as string);
    }
}
