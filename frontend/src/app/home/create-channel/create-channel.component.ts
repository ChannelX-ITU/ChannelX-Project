import { Component, OnInit, ViewEncapsulation, ViewChild } from '@angular/core';
import { Communication } from '../../models/communication';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';
import { Channel, ChannelResponse } from '../../models/channel';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { MatStepper } from '@angular/material';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
    selector: 'app-create-channel',
    templateUrl: './create-channel.component.html',
    styleUrls: ['./create-channel.component.scss']
})
export class CreateChannelComponent implements OnInit {
    @ViewChild("stepper") stepper : MatStepper;

    channel: Channel = new Channel();

    comms: Observable<Communication[]>;

    constructor(
        private store: Store<AppState>, 
        private formBuilder: FormBuilder,
        private client: HttpClient,
        private router: Router) { }

    nameGroup: FormGroup;
    commGroup: FormGroup;
    prefsGroup: FormGroup;
    ngOnInit() {
        this.comms = this.store
        .select("user")
        .pipe(map( 
          (value) => value.user.communications
          )
        );

        this.store.select("user").subscribe(value => {
            this.channel.preference = value.user.preferences;
        })

        this.nameGroup = this.formBuilder.group({
            channelName: ['', Validators.required]
        })

        this.commGroup = this.formBuilder.group({
            commValue: ['', Validators.required]
        })

        this.prefsGroup = this.formBuilder.group({})
    }

    proceed() {
        if (this.stepper.selectedIndex < 2) {
            this.stepper.selectedIndex++
        } else {
            this.client.post("/api/channels/add", {
                channel: this.channel,
                comm: this.channel.comm
            }).subscribe( () =>
                this.router.navigateByUrl("/home/channels")
            )
        }
    }
    
}
