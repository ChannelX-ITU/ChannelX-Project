import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger';

@Component({
  selector: 'app-user-preferences',
  templateUrl: './user-preferences.component.html',
  styleUrls: ['./user-preferences.component.scss']
})
export class UserPreferencesComponent implements OnInit {

  addComm: Communication = new Communication();

  comms: Observable<Communication[]>;

  constructor(
    private store: Store<AppState>,
    private client: HttpClient,
    private logger: Logger) { }

  ngOnInit() {
    this.comms = this.store
    .select("user")
    .pipe(map( 
      (value) => value.user.communications
      )
    );
  }

  addCommunication() {
    this.client.post("/api/comm/add", {
      comm_type: this.addComm.comm_type,
      value: this.addComm.value
    }).subscribe( value => {
      this.logger.log(value);
    },
    err => {
      this.logger.log(err);
    })
  }

}
