import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent, HttpErrorResponse } from '@angular/common/http'
import { Observable } from 'rxjs/Observable';
import { NotificationsService } from 'angular2-notifications'
import 'rxjs/add/operator/do';

@Injectable()
export class ErrorNotifyInterceptor implements HttpInterceptor {

    constructor(private notify: NotificationsService) {}

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
            return next.handle(req).do(event => {}, err => {
                if (err instanceof HttpErrorResponse && err.status >= 400) {
                    this.notify.error("Error", err.error["error"]["description"])
                    console.log("Error: ", err);
                }
            });
        }
}
