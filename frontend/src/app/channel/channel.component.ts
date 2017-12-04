import {ViewChild, Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-channel',
  templateUrl: './channel.component.html',
  styleUrls: ['./channel.component.scss'],
  encapsulation: ViewEncapsulation.None
})

export class ChannelComponent implements OnInit {
  constructor() { }

  regTypeSelectedOption: string = "";
  selectedNav: any;
  ngOnInit() {
      this.selectedNav = 'select value';
  }

  createNew()
  {
    console.log("Whoa");
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

  hours = [
  {value: '0100-0', view: '01:00'},
  {value: '0000-0', view: '00:00'},
  {value: '0200-0', view: '02:00'},
  {value: '0300-0', view: '03:00'},
  {value: '0400-0', view: '04:00'},
  {value: '0500-0', view: '05:00'},
  {value: '0600-0', view: '06:00'},
  {value: '0700-0', view: '07:00'},
  {value: '0800-0', view: '08:00'},
  {value: '0900-0', view: '09:00'},
  {value: '1000-0', view: '10:00'},
  {value: '1100-0', view: '11:00'},
  {value: '1200-0', view: '12:00'},
  {value: '1300-0', view: '13:00'},
  {value: '1400-0', view: '14:00'},
  {value: '1500-0', view: '15:00'},
  {value: '1600-0', view: '16:00'},
  {value: '1700-0', view: '17:00'},
  {value: '1800-0', view: '18:00'},
  {value: '1900-0', view: '19:00'},
  {value: '2000-0', view: '20:00'},
  {value: '2100-0', view: '21:00'},
  {value: '2200-0', view: '22:00'},
  {value: '2300-0', view: '23:00'},
  ];

  emails = [
    {value: '0', name: 'Default Mail', view: 'ihsanyigitergin@gmail.com'},
    {value: '1', name: 'School Mail', view: 'erginihs@itu.edu.tr'}
  ];

  phones = [
    {value: '0', name: 'Default Phone', view: '+905546544365'},
    {value: '1', name: 'School Phone', view: '+905318893565'}
  ];
}
