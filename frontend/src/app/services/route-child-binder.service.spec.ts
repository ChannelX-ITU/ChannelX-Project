import { TestBed, inject } from '@angular/core/testing';

import { RouteChildBinderService } from './route-child-binder.service';

describe('RouteChildBinderService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [RouteChildBinderService]
    });
  });

  it('should be created', inject([RouteChildBinderService], (service: RouteChildBinderService) => {
    expect(service).toBeTruthy();
  }));
});
