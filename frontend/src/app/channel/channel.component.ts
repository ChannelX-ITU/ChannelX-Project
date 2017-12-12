import {ViewChild, Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Channel, ChannelResponse } from '../models/channel';
import { Observable } from 'rxjs/Observable';
import { HttpClient } from '@angular/common/http';
import { NotificationsService } from 'angular2-notifications'
import { onErrorResumeNext, map, shareReplay } from 'rxjs/operators'
import { RouteChildBinderService } from '../services/route-child-binder.service'
import { Logger } from '@nsalaun/ng-logger';


@Component({
  selector: 'app-channel',
  templateUrl: './channel.component.html',
  styleUrls: ['./channel.component.scss'],
  providers: [RouteChildBinderService]
})

export class ChannelComponent implements OnInit {
  navLinks = [
    { path: "broadcast", label: "Broadcast Message", ownerOnly: false},
    { path: "settings", label: "Channel Settings", ownerOnly: false},
    { path: "users", label: "Channel Users", ownerOnly: true}
  ];

  channelId: string;

  currentChannel: ChannelResponse;

  constructor(
    private router: Router,
    private activatedRoute: ActivatedRoute,
    private client: HttpClient,
    private notifications: NotificationsService,
    private childBinder: RouteChildBinderService<ChannelResponse, boolean>,
    private logger: Logger) { }

  ngOnInit() {
    this.channelId = this.activatedRoute.snapshot.params["id"];
    this.logger.log("id:", this.channelId)
    
    this.client.get<ChannelResponse>("/api/channels/" + this.channelId).subscribe(value => {
      this.currentChannel = value;
      this.childBinder.notifyChild(value);
    }, error => {
      this.router.navigateByUrl("/home/channels")
    })

  }



}
