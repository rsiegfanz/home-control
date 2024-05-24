import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';

export const commonTestingProviders = [
    // Intentionally Left Blank!!!
];

export const commonTestingModules = [ReactiveFormsModule, NoopAnimationsModule, HttpClientTestingModule, RouterTestingModule] as unknown[];
