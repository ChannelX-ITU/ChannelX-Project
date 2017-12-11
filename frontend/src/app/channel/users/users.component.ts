import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { RouteChildBinderService } from '../../services/route-child-binder.service';
import { Channel, ChannelResponse } from '../../models/channel';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.scss']
})
export class UsersComponent implements OnInit {

    currentChannel: ChannelResponse = null;

  constructor(private childBinder: RouteChildBinderService<ChannelResponse, boolean>) { }

  ngOnInit() {
      this.childBinder.fromParent.subscribe( value => {
          this.currentChannel = value;
      });
  }

}
