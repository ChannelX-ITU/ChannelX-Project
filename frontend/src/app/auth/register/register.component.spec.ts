import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterComponent } from './register.component';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let debug: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [ HttpModule ],
      declarations: [ RegisterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    debug = fixture.debugElement.query(By.css('form');
    element = debug.nativeElement;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //it should get a string from Username
  it('should get a string from', () => {
    expect(element.username).toBe(typeof(text));
  });

  //it should get a string from email (maybe check for validity too?)
  it('should get a string from', () => {
    expect(element.email).toBe(typeof(text));
  });

  //it should get a string from password
  it('should get a string from', () => {
    expect(element.password).toBe(typeof(text));
  });

  //it should send a register request when register button is clicked
  it('should send register request when clicked', () => {
    let btn = element.submit;
    btn.triggerEventHandler('click', null);
    component.register().subscribe(x => {
      expect(x).toContain(true); //I don't understand this
    });
  });

});
