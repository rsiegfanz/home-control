import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';

export const routes: Routes = [
    { path: '', redirectTo: 'home', pathMatch: 'full' },
    { path: 'home', component: HomeComponent },
    {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.module').then((m) => m.DashboardModule),
    },
    {
        path: 'climateMeasurements',
        loadChildren: () => import('./climate-measurements/climate-measurements.module').then((m) => m.ClimateMeasurementsModule),
    },
    {
        path: '**',
        loadComponent: () => import('./page-not-found/page-not-found.component').then((m) => m.PageNotFoundComponent),
    },
];
