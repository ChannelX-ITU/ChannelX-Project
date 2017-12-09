import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router'
import { HttpClient } from '@angular/common/http'
import { Logger } from '@nsalaun/ng-logger'
import { User } from '../../models/user'
import { CookieService } from 'ngx-cookie'
import { Store } from '@ngrx/store'
import { UserState } from '../../state/user-state'
import { NotificationsService } from 'angular2-notifications'
import { switchMap, map } from 'rxjs/operators';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  username: string;
  password: string;

  redirectUrl: string = "/home";

  constructor(private client: HttpClient, 
    private logger: Logger, 
    private cookies: CookieService,
    private router: Router,
    private store: Store<UserState>,
    private notifications: NotificationsService,
    private activatedRoute: ActivatedRoute) { }

  ngOnInit() {
    this.activatedRoute.queryParams.subscribe(value => {
      let redirect =  value["r"]
      if (redirect) {
        this.redirectUrl = redirect;
      }
    });
    

  }

  login() {
    this.client.post("/api/login", {
      username: this.username,
      password: this.password
    }).pipe(
    switchMap( data => this.client.get<User>("/api/userinfo"))
    ).subscribe(
    data => {
      this.store.dispatch({ type: "LOGIN", user: data })
      this.router.navigateByUrl(this.redirectUrl)
    },
    error => {
    }
    )
  }

}
