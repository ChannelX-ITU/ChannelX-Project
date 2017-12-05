import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router'
import { HttpClient } from '@angular/common/http'
import { Logger } from '@nsalaun/ng-logger'
import { User } from '../../models/user'
import { CookieService } from 'ngx-cookie'
import { Store } from '@ngrx/store'
import { UserState } from '../../state/user-state'
import { NotificationsService } from 'angular2-notifications'

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
    this.client.post<User>("/api/login", {
      username: this.username,
      password: this.password
    }).subscribe(
      data => {
        this.logger.log("Got data: ", data)

        this.store.dispatch({ type: "LOGIN" })
        
        this.router.navigateByUrl("/home", {
          skipLocationChange: false
        }).then( _ => {
          this.notifications.success("Login successful!")
        })
      },
      error => {
        this.logger.error("Got error: ", error)
      }
    )
  }

}
