import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ChannelComponent } from '../channel.component';
<<<<<<< Updated upstream
import { Channel, ChannelResponse } from '../../models/channel';
import { HttpClient } from '@angular/common/http';
import { RouteChildBinderService } from '../../services/route-child-binder.service'
=======
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger'
import { Channel } from '../../models/channel';
import { RouteChildBinderService } from '../../services/route-child-binder.service';
>>>>>>> Stashed changes

@Component({
  selector: 'app-broadcast',
  templateUrl: './broadcast.component.html',
  styleUrls: ['./broadcast.component.scss']
})
export class BroadcastComponent implements OnInit {

  constructor(
      private client: HttpClient,
<<<<<<< Updated upstream
      private childBinder: RouteChildBinderService<ChannelResponse, boolean>) { }
=======
      private logger: Logger,
      private childBinder: RouteChildBinderService<Channel, boolean>) { }
>>>>>>> Stashed changes

  message: string;
  channelName: string;
  channel: Observable<Channel>;
  currentChannel: Channel = null;

  ngOnInit() {
    this.childBinder.fromParent.subscribe(value => this.channelName = value.channel.name)
  }

<<<<<<< Updated upstream
  sendMessage(){
    
=======
  SendMessage(){

    this.channel = this.childBinder.fromParent;
    this.childBinder.fromParent.subscribe(value => {
      this.logger.log("child:", value);
      this.currentChannel = value;
    });
>>>>>>> Stashed changes
    this.client.post("/api/send", {
      channel: this.currentChannel.channel.name,
      message: this.message
<<<<<<< Updated upstream
    }).subscribe();
=======
    }).subscribe(
        
    );
>>>>>>> Stashed changes
  }

}
