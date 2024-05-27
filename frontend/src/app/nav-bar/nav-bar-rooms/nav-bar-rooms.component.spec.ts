import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NavBarRoomsComponent } from './nav-bar-rooms.component';

describe('NavBarRoomsComponent', () => {
  let component: NavBarRoomsComponent;
  let fixture: ComponentFixture<NavBarRoomsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NavBarRoomsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(NavBarRoomsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
