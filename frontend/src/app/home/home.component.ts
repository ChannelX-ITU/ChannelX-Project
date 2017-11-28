import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MatTableDataSource} from '@angular/material';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  encapsulation: ViewEncapsulation.None
})

export class HomeComponent implements OnInit {
  nameForm: FormGroup;
  timeForm: FormGroup;
  commForm: FormGroup;

  displayedColumns = ['ChannelName', 'UserCount', 'CommType'];
  dataSource = new MatTableDataSource<Element>(ELEMENT_DATA);

  constructor() { }
  regTypeSelectedOption: string = "";
  selectedNav: any;
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

export interface Element {
  ChannelName: string;
  UserCount: number;
  CommType: string;
  IsActive: number;
}

const ELEMENT_DATA: Element[] = [
  {ChannelName: 'BluePanda', UserCount: 10, CommType: 'Email', IsActive: 1},
  {ChannelName: 'RedIguana', UserCount: 12, CommType: 'Email', IsActive: 1},
  {ChannelName: 'GreyChipmunk', UserCount: 23, CommType: 'Email', IsActive: 1},
  {ChannelName: 'FlyingWhale', UserCount: 2, CommType: 'SMS', IsActive: 1}
];
