import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router'
import { Store } from '@ngrx/store'
import { Observable } from 'rxjs/Observable';
import { Logger } from '@nsalaun/ng-logger'
import { AppState } from '../state/app-state'


@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {

  logged_in : boolean

  constructor(private store: Store<AppState>, private router: Router, private logger: Logger) { }

  ngOnInit() {
      this.logger.info(this.store);
      this.store.select('user').subscribe(data => {
          this.logged_in = data.logged_in
          this.logger.info(data)
      })
  }

  logout() {
      this.store.dispatch({ type: "LOGOUT" })
      this.router.navigateByUrl("/login")
  }

}
