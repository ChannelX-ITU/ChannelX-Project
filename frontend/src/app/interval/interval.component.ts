import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { Interval } from '../models/interval'
import { Logger } from '@nsalaun/ng-logger';

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

    constructor(private logger: Logger) { }

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
        this.start = this.closestMultiple((interval.start % 1440), 15);
        this.end = this.closestMultiple((interval.start + interval.duration) % 1440, 15);
        this.days.push(((interval.start / 1440) | 0) % 7); //js hack to convert float to int
    }

    closestMultiple(value: number, multiple: number) : number {
        return ((value / multiple) | 0) * multiple;
    }

    toIntervals() : Interval[] {
        return this.days.map((value) => {
            let start = value * 1440 + this.start;
            let duration = value * 1440 + this.end - start;
            return new Interval(start, duration);
        })
    }
}
