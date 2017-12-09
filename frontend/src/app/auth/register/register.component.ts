import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { NotificationsService } from 'angular2-notifications';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

    username: string;
    password: string;
    email: string;

    constructor(
        private client: HttpClient, 
        private notifications: NotificationsService,
        private router: Router) { }

    ngOnInit() {
    }

    register() {
        this.client.post("/api/signup", {
            username: this.username,
            password: this.password,
            email: this.email
        }).subscribe(
            resp => {
                this.router.navigateByUrl("/login").then(_ =>{
                });
            }
        )
    }

}
