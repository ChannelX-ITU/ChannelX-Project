import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ChannelComponent } from '../channel.component';
import { Channel, ChannelResponse } from '../../models/channel';
import { HttpClient } from '@angular/common/http';
import { RouteChildBinderService } from '../../services/route-child-binder.service'

@Component({
  selector: 'app-broadcast',
  templateUrl: './broadcast.component.html',
  styleUrls: ['./broadcast.component.scss']
})
export class BroadcastComponent implements OnInit {

  constructor(
      private client: HttpClient,
      private childBinder: RouteChildBinderService<ChannelResponse, boolean>) { }

  message: string;
  channelName: string;

  ngOnInit() {
    this.childBinder.fromParent.subscribe(value => this.channelName = value.channel.name)
  }

  sendMessage(){
    
    this.client.post("/api/send", {
      channel: this.channelName,
      message: this.message
    }).subscribe();
  }

}
