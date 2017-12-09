import {ViewChild, Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ChannelResponse } from '../models/channel';
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
    { path: "broadcast", label: "Broadcast Message"},
    { path: "settings", label: "Channel Settings"}
  ];

  channelId: string;

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
      this.childBinder.notifyChild(value);
      this.logger.log("channel comp: ", value);
    }, error => {
      this.logger.log("channel err: ", error);
      this.router.navigateByUrl("/home/channels").then(_ => {
        this.notifications.error("Error", error.message)
      })
    })

  }

  

}
