import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router'
import { Store } from '@ngrx/store'
import { Observable } from 'rxjs/Observable';
import { Logger } from '@nsalaun/ng-logger'
import { AppState } from '../state/app-state'
import { User } from '../models/user';
import { switchMap, map } from 'rxjs/operators';


@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

  user : Observable<User>
  logged_in : Observable<boolean>;

  constructor(private store: Store<AppState>, private router: Router, private logger: Logger) { }

  ngOnInit() {
      this.logger.info(this.store);
      this.store.select('user').subscribe( data => {
        this.logger.log("Subscribe: " , data);
      })
      this.user = this.store.select('user').pipe(map( (data, index) => {
        this.logger.log("USER: " , data);
        return data.user;
      }));
      this.logged_in = this.store.select('user').pipe(map( (data, index) => {
        this.logger.log("LOGGED_IN: " , data);
        return data.logged_in;
      }))
  }

  logout() {
      this.store.dispatch({ type: "LOGOUT" })
      this.router.navigateByUrl("/login")
  }

}
