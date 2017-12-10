import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'
import { Restriction } from '../../models/restriction'
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger';
import { User } from '../../models/user';

@Component({
  selector: 'app-user-preferences',
  templateUrl: './user-preferences.component.html',
  styleUrls: ['./user-preferences.component.scss']
})
export class UserPreferencesComponent implements OnInit {

  addComm: Communication = new Communication();

  addRest: Restriction = new Restriction();

  user: User = new User();

  constructor(
    private store: Store<AppState>,
    private client: HttpClient,
    private logger: Logger) { }

  ngOnInit() {
    this.store
    .select("user")
    .subscribe(val => {this.user = val.user; console.log(val.user)});
  }

  removeCommunication(c: Communication) {
    let comm = new Communication(c.comm_type, c.value)

    comm.remove(this.client).subscribe( user => {
      this.store.dispatch({type: "REFRESH", user: user})
    })
  }

  addCommunication() {
    this.addComm.add(this.client).subscribe( user => {
      this.store.dispatch({type: "REFRESH", user: user})
    })
  }

  addRestriction(){
    
  }

  save() {
    this.client.post("/api/userinfo/update", this.user.preferences).subscribe();
  }

}
