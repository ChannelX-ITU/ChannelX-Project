import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router'
import { Logger } from '@nsalaun/ng-logger'

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class AuthComponent implements OnInit {

    navLinks = [
      { path: "/login", label: "Login"},
      { path: "/register", label: "Register"}
    ];

  constructor(private route: ActivatedRoute, private router: Router, private _logger: Logger) { 
      
  }

  ngOnInit() {
  }

}
