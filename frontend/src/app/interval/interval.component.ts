import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { Interval } from '../models/interval'

@Component({
    selector: 'app-interval',
    templateUrl: './interval.component.html',
    styleUrls: ['./interval.component.scss'],
    encapsulation: ViewEncapsulation.None
})

export class IntervalComponent implements OnInit {

    @Input() 
    intervals: IntervalInterface[];

    @Output()
    intervalsChange: EventEmitter<IntervalInterface[]> = new EventEmitter<IntervalInterface[]>();

    constructor() { }

    ngOnInit() {
    }

    close(interval: IntervalInterface) {
        this.intervals = this.intervals.filter( (value) => {
            return value != interval;
        });
        this.intervalsChange.emit(this.intervals);
    }

}


export class IntervalInterface {
    days: number[] = [];
    start: number = 0;
    end: number = 0;

    constructor(interval: Interval = null) {
        if (interval === null) return;
        this.start = interval.start % 1440;
        this.end = (interval.start + interval.duration) % 1440;
        this.days.push(((interval.start / 1440) | 0) % 7); //js hack to convert float to int
    }
}
