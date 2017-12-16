import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement }    from '@angular/core';
import { By }              from '@angular/platform-browser';
import { AuthComponent } from './auth.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

describe('AuthComponent', () => {
  let component: AuthComponent;
  let fixture: ComponentFixture<AuthComponent>;
  let debug: DebugElement;
  let element: HTMLElement;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AuthComponent, LoginComponent, RegisterComponent, RouterLinkStubDirective, RouterOutletStubComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthComponent); //this is the test subject
    component = fixture.componentInstance;

    debug = fixture.debugElement.query(By.css('router-outlet'); //I want to look at the router outlet which I understand displays login/register comps
    element = debug.nativeElement;
    fixture.detectChanges();

    linkDes = fixture.debugElement
      .queryAll(By.directive(RouterLinkStubDirective));

    // get the attached link directive instances using the DebugElement injectors
    links = linkDes
      .map(de => debugElement.injector.get(RouterLinkStubDirective) as RouterLinkStubDirective);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  //it should display login component
  it('should display', () => {
    expect(element).toContain(LoginComponent);
  });

  //it can get router links and confirm their validity
  it('can get router links', () => {
      expect(links.length).toBe(2, 'should have 2 links');
      expect(links[0].linkParams).toBe('/login', '1st link should go to login');
      expect(links[1].linkParams).toBe('/register', '2nd link should go to register');
  });

  //it can click on register link
  it('can click on register', () => {
    const registerLinkDe = linkDes[1];
    const regsiterLink = links[1];

    expect(regsiterLink.navigatedTo).toBeNull('link should not have navigated yet');

    registerLinkDe.triggerEventHandler('click', null); //this simulates the click event
    fixture.detectChanges();

    expect(regsiterLink.navigatedTo).toBe('/register');
  });

  //it should display register component (with detectChanges)
  it('should display', () => {
    fixture.detectChanges();
    expect(element).toContain(RegisterComponent);
  });
});
