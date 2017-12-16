import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LoginComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //it should get a string from Username
  it('should get a string from', () => {
    expect(element.username).toBe(typeof(text));
  });
  
  //it should get a string from password
  it('should get a string from', () => {
    expect(element.password).toBe(typeof(text));
  });

  //it should send a login request when login button is clicked

  //it should route to home component if login request returns true (assume returns true)

});
