import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router'
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

  constructor(private client: HttpClient, 
    private logger: Logger, 
    private cookies: CookieService,
    private router: Router,
    private store: Store<UserState>,
    private notifications: NotificationsService) { }

  ngOnInit() {
    this.cookies.put("test", "yiha");
  }

  login() {
    this.cookies.get("test")
    this.client.post("/api/login", {
      username: this.username,
      password: this.password
    }).pipe(switchMap( data => {
      this.logger.log("Got data @ map: ", data);
      return this.client.get<User>("/api/userinfo");
    })).subscribe(
      data => {
        this.store.dispatch({ type: "LOGIN", user: data })
        this.router.navigateByUrl("/home")
      },
      error => {
      }
    )
  }

}
