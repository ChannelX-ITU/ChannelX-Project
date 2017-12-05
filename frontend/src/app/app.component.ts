import { Component } from '@angular/core';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {
    public options = {
        position: ["bottom", "right"],
        timeOut: 3000,
        showProgressBar: false,
        lastOnBottom: true,
        theClass: "notification"
    }
    title = 'app';
}
