import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { SuccessMessage } from '../models/reponses';
import { switchMap } from 'rxjs/operators';
import { User } from '../models/user';

export class Communication {
    constructor(
        public comm_type: string = "EMAIL",
        public value: string = ""
        ) {}

    public log(val: string) {
        console.log(val);
    }

    public static add(client: HttpClient, comm: Communication): Observable<User> {
        return Communication.post(client, "/api/comm/add", comm)

    }
    public static remove(client: HttpClient, comm: Communication): Observable<User> {
        return Communication.post(client, "/api/comm/remove", comm)
    }

    public add(client: HttpClient): Observable<User> {
        return Communication.add(client, this);
    }
    
    public remove(client: HttpClient): Observable<User> {
        return Communication.remove(client, this);
    }

    private static post(client: HttpClient, url: string, comm: Communication): Observable<User> {
        return client.post(url, {
            comm_type: comm.comm_type,
            value: comm.value
        }).pipe(
        switchMap( () => client.get<User>("/api/userinfo") )
        )
    }
}
