import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-reply',
    templateUrl: './reply.component.html',
    styleUrls: ['./reply.component.scss']
})
export class ReplyComponent implements OnInit {

    message: string;
    token: string;

    constructor(private client: HttpClient, private activatedRoute: ActivatedRoute, private router: Router) { }

    ngOnInit() {
        this.activatedRoute.queryParams.subscribe(params => {
            this.token = params["t"];
        })
    }

    send() {
        this.client.post("/api/send/" + this.token, {
            message: this.message
        }).subscribe( () => {
            this.router.navigateByUrl("/home");
        })
    }

}
