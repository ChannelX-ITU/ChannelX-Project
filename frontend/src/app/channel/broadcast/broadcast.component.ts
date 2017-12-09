import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ChannelComponent } from '../channel.component';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-broadcast',
  templateUrl: './broadcast.component.html',
  styleUrls: ['./broadcast.component.scss']
})
export class BroadcastComponent implements OnInit {

  constructor(
      private client: HttpClient) { }

  message: string;
  channelName: string;

  ngOnInit() {
  }

  SendMessage(){
    
    this.channelName = document.getElementById('channelName').innerHTML;
    this.client.post("/api/send", {
      channel: this.channelName,
      message: this.message
    }).subscribe(

    );
  }

}
