import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardRoomsComponent } from './dashboard-rooms.component';

describe('DashboardRoomsComponent', () => {
    let component: DashboardRoomsComponent;
    let fixture: ComponentFixture<DashboardRoomsComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [DashboardRoomsComponent],
        }).compileComponents();

        fixture = TestBed.createComponent(DashboardRoomsComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});

