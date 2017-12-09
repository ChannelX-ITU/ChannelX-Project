import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { IntervalInterface } from '../interval/interval.component'
import { Preference } from '../models/preference';

@Component({
    selector: 'app-preferences',
    templateUrl: './preferences.component.html',
    styleUrls: ['./preferences.component.scss']
})
export class PreferencesComponent implements OnInit {

    @Input()
    preference: Preference;

    @Output()
    preferenceChange: EventEmitter<Preference> = new EventEmitter<Preference>();

    constructor() { }

    ngOnInit() {
        
    }

}
