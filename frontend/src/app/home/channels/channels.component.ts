import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import {MatTableDataSource} from '@angular/material';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map, shareReplay } from 'rxjs/operators';
import { Communication } from '../../models/communication';
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger';

interface UserChannels {
  owned: ChannelInterface[];
  subbed: ChannelInterface[];
}

interface ChannelInterface {
  name: string;
  user_count: number;
  comm: string;
}


@Component({
  selector: 'app-channels',
  templateUrl: './channels.component.html',
  styleUrls: ['./channels.component.scss']
})
export class ChannelsComponent implements OnInit {


  displayedColumns = ['ChannelName', 'UserCount', 'CommType'];
  ownedDataSource = new MatTableDataSource<ChannelInterface>();
  subscribedDataSource = new MatTableDataSource<ChannelInterface>();
  comms: Observable<Communication[]>;
  channels: UserChannels;

  accordions = [
    { source: this.ownedDataSource , title: "Owned Channels"},
    { source: this.subscribedDataSource, title: "Subscribed Channels"}
  ]

  loaded = false;

  openedAcc: number = -1;

  channelName: string;
  comm: string;

  constructor(private store: Store<AppState>, private client: HttpClient, private logger: Logger) { }

  ngOnInit() {
    this.updateChannels().subscribe(() => {
      this.openAccordion(0);
    })
    this.comms = this.store
    .select("user")
    .pipe(map(
      (value) => value.user ? value.user.communications : []
      )
    );
  }

  updateChannels() : Observable<UserChannels> {
    this.loaded = false;

    let obs = this.client.get<UserChannels>("/api/channels").pipe(shareReplay(3));
    obs.subscribe( data => {
      this.loaded = true;
      this.channels = data;
      this.ownedDataSource.data = this.channels.owned;
      this.subscribedDataSource.data = this.channels.subbed;
    });
    return obs;
  }

  openAccordion(accordion: number) {
    // console.log("WUT");
    this.openedAcc = accordion;
  }

  joinChannel()
  {
    this.client.post("/api/channels/join", {
      channel: this.channelName,
      comm: this.comm
    }).subscribe( () => {
      this.channelName = "";
      this.comm = "";
      this.updateChannels().subscribe(() => {
        this.openAccordion(1);
      })
    });


  }

}
