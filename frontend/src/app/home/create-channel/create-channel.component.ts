import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
    selector: 'app-create-channel',
    templateUrl: './create-channel.component.html',
    styleUrls: ['./create-channel.component.scss']
})
export class CreateChannelComponent implements OnInit {


    regTypeSelectedOption: string = "";
    selectedNav: any;

    constructor() { }

    ngOnInit() {
        this.selectedNav = 'select value';

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
