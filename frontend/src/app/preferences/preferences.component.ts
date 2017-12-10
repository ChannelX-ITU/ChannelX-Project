import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { IntervalInterface } from '../interval/interval.component'
import { Preference } from '../models/preference';
import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';
import { merge, filter } from 'rxjs/operators';
import {MatDatepickerInputEvent} from '@angular/material/datepicker';

@Component({
    selector: 'app-preferences',
    templateUrl: './preferences.component.html',
    styleUrls: ['./preferences.component.scss']
})
export class PreferencesComponent implements OnInit {

    @Input()
    readOnly: boolean = false;

    @Input()
    preference: Preference;

    @Output()
    preferenceChange: EventEmitter<Preference> = new EventEmitter<Preference>();

    preferenceInterface: PreferenceInterface;

    constructor() { }

    ngOnInit() {
        console.log(this.readOnly);
        this.preferenceInterface = new PreferenceInterface(this.preference);
        this.preferenceInterface.onChange.asObservable().pipe(filter(val => val)).subscribe(() => {
            this.updatePreference()
        });
        this.updatePreference();
    }

    updatePreference() {
        let vals = this.preferenceInterface.toPreferenceValues();
        console.log("updated preference dates", vals);
        this.preference.start_date = vals["start_date"];
        this.preference.duration = vals["duration"];
        if (this.preference.duration == 0) {
            this.preference.start_date = 0;
        }
    }

    dateChange(event: MatDatepickerInputEvent<Date>) {
        console.log(event);
    }

}

export class PreferenceInterface {
    start: Date;
    end: Date;

    min: Date = new Date()

    onChange: Subject<boolean> = new Subject<boolean>()

    constructor(preference: Preference) {
        this.start = new Date(preference.start_date);
        this.end = new Date(preference.start_date + preference.duration*86400000);
    }

    updateValue(newValue: any) {
        this.onChange.next(true);
    }
    
    createDateAsUTC(date) {
        return new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
    }

    toPreferenceValues(): Object {
        console.log(this);
        let start = this.createDateAsUTC(this.start).valueOf()
        return {
            start_date: start,
            duration: ((this.createDateAsUTC(this.end).valueOf()-start)/86400000) | 0

        }
    }

}
