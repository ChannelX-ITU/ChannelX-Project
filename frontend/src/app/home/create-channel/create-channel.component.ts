import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Communication } from '../../models/communication';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';
import { map } from 'rxjs/operators';

@Component({
    selector: 'app-create-channel',
    templateUrl: './create-channel.component.html',
    styleUrls: ['./create-channel.component.scss']
})
export class CreateChannelComponent implements OnInit {


    intervals: Communication[] = [];

    comms: Observable<Communication[]>;
    regTypeSelectedOption: string = "";
    selectedNav: any;

    constructor(private store: Store<AppState>) { }

    ngOnInit() {
        this.selectedNav = 'select value';
        this.comms = this.store
        .select("user")
        .pipe(map( 
          (value) => value.user.communications
          )
        );
    }
    
    setNav(nav:any){
        this.selectedNav = nav;
        if(this.selectedNav == "Email"){
            this.regTypeSelectedOption = "email";
        }
        else if(this.selectedNav == "SMS"){
            this.regTypeSelectedOption = "sms";
        }
    }

    onSubmit(username: String) {
        console.log(username);
    }


    channels = [
    {value: "ChannelOne-0", view: "BluePanda"},
    {value: "ChannelTwo-1", view: "VignetteTiger"}
    ]

    commTypes = [
    {value: 'Email', view: 'Email'},
    {value: 'SMS', view: 'SMS'}
    ];

    days = [
    {value: 'Monday-0', view: 'Monday'},
    {value: 'Tuesday-1', view: 'Tuesday'},
    {value: 'Wednesday-2', view: 'Wednesday'},
    {value: 'Thursday-3', view: 'Thursday'},
    {value: 'Friday-4', view: 'Friday'},
    {value: 'Saturday-5', view: 'Saturday'},
    {value: 'Sunday-6', view: 'Sunday'}
    ];
}
