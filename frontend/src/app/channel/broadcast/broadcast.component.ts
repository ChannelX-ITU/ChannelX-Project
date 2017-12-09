import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ChannelComponent } from '../channel.component';
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger'
import { Channel } from '../../models/channel';
import { RouteChildBinderService } from '../../services/route-child-binder.service';

@Component({
  selector: 'app-broadcast',
  templateUrl: './broadcast.component.html',
  styleUrls: ['./broadcast.component.scss']
})
export class BroadcastComponent implements OnInit {

  constructor(
      private client: HttpClient,
      private logger: Logger,
      private childBinder: RouteChildBinderService<Channel, boolean>) { }

  message: string;
  channelName: string;
  channel: Observable<Channel>;
  currentChannel: Channel = null;

  ngOnInit() {
    this.childBinder.fromParent.subscribe(value => this.channelName = value.name)
  }

  SendMessage(){

    this.channel = this.childBinder.fromParent;
    this.childBinder.fromParent.subscribe(value => {
      this.logger.log("child:", value);
      this.currentChannel = value;
    });
    this.client.post("/api/send", {
      channel: this.currentChannel.name,
      message: this.message
    }).subscribe(

    );
  }

}
