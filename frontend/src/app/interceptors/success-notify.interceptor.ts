import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent, HttpErrorResponse, HttpResponse } from '@angular/common/http'
import { Observable } from 'rxjs/Observable';
import { NotificationsService } from 'angular2-notifications'
import 'rxjs/add/operator/do';

@Injectable()
export class SuccessNotifyInterceptor implements HttpInterceptor {

    constructor(private notify: NotificationsService) {}

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
            return next.handle(req).do(event => {
                if (event instanceof HttpResponse && event.body && event.body.message) {
                    this.notify.success("Success", event.body.message)

                }
            }, err => {});
        }
}
