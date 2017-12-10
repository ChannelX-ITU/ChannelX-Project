import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { IndexerPipe } from '../../pipes/indexer.pipe';
import { IntervalInterface } from '../../interval/interval.component'
import { Interval } from '../../models/interval';
import { Logger } from '@nsalaun/ng-logger'
import { Channel, ChannelResponse } from '../../models/channel';
import { Observable } from 'rxjs/Observable';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'
import { RouteChildBinderService } from '../../services/route-child-binder.service'
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';


@Component({
  selector: 'app-edit-channel',
  templateUrl: './edit-channel.component.html',
  styleUrls: ['./edit-channel.component.scss']
})
export class EditChannelComponent implements OnInit {


  channel: Observable<Channel>;

  comms: Observable<Communication[]>;

  currentChannel: ChannelResponse = null;

  constructor(
    private logger: Logger, 
    private store: Store<AppState>, 
    private childBinder: RouteChildBinderService<ChannelResponse, boolean>,
    private client: HttpClient,
    private router: Router) { }

  ngOnInit() {
    this.channel = this.childBinder.fromParent.map( val => val.channel);
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
    this.client.post("/api/channels/" + this.currentChannel.channel.name + "/update", 
      {
        alias: this.currentChannel.alias,
        channel: this.currentChannel.channel,
        comm: this.currentChannel.comm.value
      }).subscribe();
  }

  destroy() {
    this.client.post("/api/channels/leave", {
      channel: this.currentChannel.channel.name
    }).subscribe(() => {
      this.router.navigateByUrl("/home");
    });
  }

}
