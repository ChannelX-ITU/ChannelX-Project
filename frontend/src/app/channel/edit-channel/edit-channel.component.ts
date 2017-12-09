import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { IndexerPipe } from '../../pipes/indexer.pipe';
import { IntervalInterface } from '../../interval/interval.component'
import { Interval } from '../../models/interval';
import { Logger } from '@nsalaun/ng-logger'
import { Channel } from '../../models/channel';
import { Observable } from 'rxjs/Observable';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'
import { RouteChildBinderService } from '../../services/route-child-binder.service'


@Component({
  selector: 'app-edit-channel',
  templateUrl: './edit-channel.component.html',
  styleUrls: ['./edit-channel.component.scss']
})
export class EditChannelComponent implements OnInit {


  channel: Observable<Channel>;

  comms: Observable<Communication[]>;

  currentChannel: Channel = null;

  constructor(
    private logger: Logger, 
    private store: Store<AppState>, 
    private childBinder: RouteChildBinderService<Channel, boolean>) { }

  ngOnInit() {
    this.channel = this.childBinder.fromParent;
    this.childBinder.fromParent.subscribe(value => {
      this.logger.log("child:", value);
      this.currentChannel = value;
    });
    this.comms = this.store
    .select("user")
    .pipe(map( 
      (value) => value.user ? value.user.communications.sort((a,b) => (a.comm_type > b.comm_type) ? 1 : -1) : new Array<Communication>()
      )
    );
  }

  save() {
    // let intervals = this.intervals.map( value => value.toIntervals()).reduce((acc, value) => {
    //   value.forEach( value => acc.push(value));
    //   return acc;
    // }, []);
    this.logger.log("Current channel: " , this.currentChannel);
  }

}
