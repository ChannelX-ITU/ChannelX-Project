import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { IntervalInterface } from '../interval/interval.component'
import { Preference } from '../models/preference';
import { Restriction, RestrictionType } from '../models/restriction';
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

    restType = Restriction;

    @Input()
    showRestrictions: boolean = false;

    addRest: Restriction = new Restriction();

    @Input()
    restrictions: Restriction[];

    @Output()
    restrictionsChange: EventEmitter<Restriction[]> = new EventEmitter<Restriction[]>();

    @Input()
    readOnly: boolean = false;

    @Input()
    preference: Preference;

    @Output()
    preferenceChange: EventEmitter<Preference> = new EventEmitter<Preference>();

    preferenceInterface: PreferenceInterface;

    noExpiration: boolean = false;

    constructor() { }

    ngOnInit() {
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
        this.preference.duration = vals["duration"] + 1;
        if (this.noExpiration) {
            this.preference.duration = 0;
        }
    }

    dateChange(event: MatDatepickerInputEvent<Date>) {
        console.log(event);
    }

    addRestriction() {
        this.restrictions.push(this.addRest);
        this.addRest = new Restriction();
        this.restrictionsChange.next(this.restrictions);
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
    
    createDateAsUTC(date: Date) {
        return date;
    }

    toPreferenceValues(): Object {
        console.log(this);
        let start = this.start.valueOf()
        return {
            start_date: start,
            duration: ((this.end.valueOf()-start)/86400000) | 0

        }
    }

}
