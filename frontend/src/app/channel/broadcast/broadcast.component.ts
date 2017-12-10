import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ChannelComponent } from '../channel.component';
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';
import { Logger } from '@nsalaun/ng-logger'
import { Channel, ChannelResponse } from '../../models/channel';
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
      private childBinder: RouteChildBinderService<ChannelResponse, boolean>) { }

  message: string;
  channel: Observable<Channel>;
  currentChannel: ChannelResponse = null;

  ngOnInit() {
    this.childBinder.fromParent.subscribe(value => {
      this.logger.log("child:", value);
      this.currentChannel = value;
    });
  }

  sendMessage(){
    let msg = this.message;
    msg += `
    <html>
    <head>
      <style>
        .colored {
          color: blue;
        }
        #body {
          font-size: 14px;
        }
      </style>
    </head>
    <body>
      <div id='body'>
      <span>You can use this snippet to answer to ` + this.currentChannel.channel.name + ` owner</span>
        <form method="POST" action="http://localhost:6969/api/send">
          <input type="text" placeholder="Type your response here.."><br/>
          <button class='colored'>SEND</button>
        </form>
      </div>
    </body>
    </html>
    `

    this.client.post("/api/send", {
      channel: this.currentChannel.channel.name,
      message: this.message
    }).subscribe(

    );
  }

}
