import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter, forwardRef } from '@angular/core';
import { Communication } from '../models/communication';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
    selector: 'app-communications',
    templateUrl: './communications.component.html',
    styleUrls: ['./communications.component.scss'],
    providers: [{ 
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => CommunicationsComponent),
      multi: true
    }]
})
export class CommunicationsComponent implements OnInit, ControlValueAccessor {

    @Input()
    communications: Communication[];

    selection: string;

    propagateChange = (_: any) => {};
    
    constructor() { }

    ngOnInit() {
    }

    updateValue(newValue: string) {
        this.selection = newValue;
        this.propagateChange(this.selection);
    }

    writeValue(value: string) {
        this.selection = value;
    }

    registerOnChange(fn) {
        this.propagateChange = fn;
    }

    registerOnTouched() {}

}
