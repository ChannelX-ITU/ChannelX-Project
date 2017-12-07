import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import { Store } from '@ngrx/store';
import { AppState } from '../state/app-state';
import { CookieService } from 'ngx-cookie';
import { of } from 'rxjs/observable/of';
import { map, onErrorResumeNext } from 'rxjs/operators';
import { Logger } from '@nsalaun/ng-logger';
import { User } from '../models/user';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(private store: Store<AppState>, 
        private router: Router, 
        private cookies: CookieService, 
        private client: HttpClient,
        private logger: Logger) {}

    canActivate(
        next: ActivatedRouteSnapshot,
        state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {

        let sessionToken = this.cookies.get("bist-sissin-ivir");
        this.logger.log("token: ", sessionToken);

        if (sessionToken != "") {
            return this.client.get<User>("/api/userinfo").pipe(onErrorResumeNext(of(null))).map((value) => {
                if (value === null) {
                    this.router.navigateByUrl("/login");
                } else {
                    this.store.dispatch({type: "AUTO_LOGIN", user: value});
                }
                return value !== null;

            })
        }
        return false;
        // return this.store.select("user").pipe().map( user => {
        //     if (!user.logged_in) {
        //         this.router.navigateByUrl("/login");
        //     }
        //     return user.logged_in;
        // }, this);
    }
}
