import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { map, onErrorResumeNext } from 'rxjs/operators';

@Injectable()
export class ActivationGuard implements CanActivate {
    constructor(
        private router: Router, 
        private client: HttpClient
        ) {}
  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
      let token = next.params["token"];
    return this.client.get("/api/activate/" + token).pipe(onErrorResumeNext(of(null))).map(() => {
        this.router.navigateByUrl("/login");
        return false;
    })
  }
}
