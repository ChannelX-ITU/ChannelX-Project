import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {MatTableDataSource} from '@angular/material';

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

  constructor() { }
  ngOnInit() {
  }



}

