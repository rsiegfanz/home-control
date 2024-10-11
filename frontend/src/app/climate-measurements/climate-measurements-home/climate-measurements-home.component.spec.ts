import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClimateMeasurementsHomeComponent } from './climate-measurements-home.component';

describe('ClimateMeasurementsHomeComponent', () => {
  let component: ClimateMeasurementsHomeComponent;
  let fixture: ComponentFixture<ClimateMeasurementsHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClimateMeasurementsHomeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClimateMeasurementsHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
