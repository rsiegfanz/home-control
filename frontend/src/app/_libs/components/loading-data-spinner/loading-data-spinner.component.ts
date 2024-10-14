import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

@Component({
    selector: 'app-loading-data-spinner',
    standalone: true,
    imports: [CommonModule, ProgressSpinnerModule],
    templateUrl: './loading-data-spinner.component.html',
    styleUrl: './loading-data-spinner.component.scss',
})
export class LoadingDataSpinnerComponent {
    @Input() loadingText: string = '';
    @Input() errorText: string = '';
}
