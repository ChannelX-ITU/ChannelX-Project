import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'

@Component({
  selector: 'app-user-preferences',
  templateUrl: './user-preferences.component.html',
  styleUrls: ['./user-preferences.component.scss']
})
export class UserPreferencesComponent implements OnInit {

  comms: Observable<Communication[]>;

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
    this.comms = this.store
    .select("user")
    .pipe(map( 
      (value) => value.user.communications
      )
    );
  }

}
