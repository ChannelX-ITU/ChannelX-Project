import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
import { shareReplay } from 'rxjs/operators';

@Injectable()
export class RouteChildBinderService<T, R> {

  private fromParentSubject: Subject<T> = new Subject();
  private toParentSubject: Subject<R> = new Subject();

  fromParent: Observable<T> = this.fromParentSubject.asObservable().pipe(shareReplay(3))
  toParent: Observable<R> = this.toParentSubject.asObservable().pipe(shareReplay(3))


  constructor() { 
      this.fromParent.subscribe();
      this.toParent.subscribe();
  }

  notifyChild(value: T) {
      this.fromParentSubject.next(value);
  }

  notifyParent(value: R) {
      this.toParentSubject.next(value);
  }

}
