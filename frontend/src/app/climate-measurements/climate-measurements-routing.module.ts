import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClimateMeasurementsHomeComponent } from './climate-measurements-home/climate-measurements-home.component';

const routes: Routes = [{ path: '', component: ClimateMeasurementsHomeComponent }];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class ClimateMeasurementsRoutingModule {}
