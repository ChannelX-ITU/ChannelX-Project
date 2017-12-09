import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import {MatTableDataSource} from '@angular/material';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
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
  loaded = false;

  channelName: string;
  comm: string;

  constructor(private store: Store<AppState>, private client: HttpClient, private logger: Logger) { }

  ngOnInit() {
    this.client.get<UserChannels>("/api/channels").subscribe( data => {
      this.ownedDataSource.data = data.owned;
      this.subscribedDataSource.data = data.subbed;
      this.loaded = true;
      this.logger.log("Channels:", data);
    });
    this.comms = this.store
    .select("user")
    .pipe(map(
      (value) => value.user.communications
      )
    );
  }

  joinChannel()
  {
    this.client.post("/api/channels/join", {
      channel: this.channelName,
      message: this.comm
    }).subscribe(

    );
  }

}

export interface Element {
  ChannelName: string;
  UserCount: number;
  CommType: string;
  IsActive: number;
}

const ELEMENT_DATA: Element[] = [
{ChannelName: 'BluePanda', UserCount: 10, CommType: 'Email', IsActive: 1},
{ChannelName: 'RedIguana', UserCount: 12, CommType: 'Email', IsActive: 1},
{ChannelName: 'GreyChipmunk', UserCount: 23, CommType: 'Email', IsActive: 1},
{ChannelName: 'FlyingWhale', UserCount: 2, CommType: 'SMS', IsActive: 1}
];
