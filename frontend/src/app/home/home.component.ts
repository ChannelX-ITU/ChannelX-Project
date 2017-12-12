import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MatTableDataSource} from '@angular/material';
import { HttpClient } from '@angular/common/http';
import { Channel } from '../models/channel';
import { Logger } from '@nsalaun/ng-logger';
import { User } from '../models/user';
import { AppState } from '../state/app-state';
import { Store } from '@ngrx/store';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit {


  navLinks = [
  { path: "channels", label: "Channels"},
  { path: "preferences", label: "User preferences"},
  { path: "create", label: "Create Channel"}
  ];

  constructor(
    private client: HttpClient, 
    private logger: Logger, 
    private store: Store<AppState>
    ) { }

  ngOnInit() {
    this.logger.log("init");
    // this.store.select("user").pipe(map( value => {
    //   this.logger.log("MAP: ", value)
    //   return value;
    // }),skipWhile( (value) => {
    //   this.logger.log("SKIP: ", value, value === null)
    //   return value === null
    // }), 
    //   switchMap( (value) => {
    //     this.logger.log("SWITCH: ", value)

    //     let channels = value.user.channels.map(
    //       (channel_name) => this.client.get<Channel>("/api/channels/" + channel_name)
    //       ).reduce( (prev, curr) => {return prev.pipe(merge(curr))}, of(new Channel()))
    //     this.logger.log("SWITCH channels: ", channels)

    //     return channels;
    //     }),
    //   reduce( (acc: Channel[], val: Channel) => {
    //     acc.push(val);
    //     this.logger.log(acc);
    //     return acc;
    //   }, Array())
    //   ).subscribe( channels => this.logger.log("Channels: ", channels));
  }

}


