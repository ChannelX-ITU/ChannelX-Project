import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import { Interval } from '../models/interval'
import { Logger } from '@nsalaun/ng-logger';
import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { merge, filter } from 'rxjs/operators';

@Component({
    selector: 'app-interval',
    templateUrl: './interval.component.html',
    styleUrls: ['./interval.component.scss']
})

export class IntervalComponent implements OnInit {

    @Input()
    readOnly: boolean = false;

    @Input()
    data: Interval[];

    @Output()
    dataChange: EventEmitter<Interval[]> = new EventEmitter<Interval[]>();

    intervals: IntervalInterface[] = [];

    constructor(private logger: Logger) { }

    ngOnInit() {
        IntervalInterface.convertIntervals(this.intervals, this.data);

        // console.log("this.intervals=", this.intervals);
        this.intervals.forEach( val => {
            val.onChange.asObservable().subscribe( _ => {
                this.updateIntervals()
            })
        })
    }

    updateIntervals() {
        console.log("Updating intervals...");
        this.data = IntervalInterface.toIntervals(this.intervals)
        console.log("this.data=", this.data);
        this.dataChange.next(this.data);
    }

    close(interval: IntervalInterface) {
        this.intervals.forEach( (value, index) => {
            let res = value != interval;
            if (!res) {
                value.onChange.unsubscribe();
                this.intervals.splice(index, 1);
            }
        });
        this.updateIntervals();
    }

    createNew()
    {
        let intrfc = new IntervalInterface();
        intrfc.onChange.asObservable().subscribe(_ => {
            this.updateIntervals();
        })
        // this.intervalWatcher.pipe(merge(intrfc.onChange.asObservable()))
        this.intervals.push(intrfc)
        this.logger.log(this.intervals);
    }
}


export class IntervalInterface {
    days: number[] = [];
    start: number = 0;
    end: number = 0;

    onChange: Subject<boolean> = new Subject<boolean>()

    constructor(interval: Interval = null) {
        if (interval === null) return;
        let date = new Date();
        let start = this.fromUTC(interval.start);
        let end = start + interval.duration;
        this.start = this.closestMultiple((start % 1440), 15);
        this.end = this.closestMultiple(end % 1440, 15);
        this.days.push(((start / 1440) | 0) % 7); //js hack to convert float to int
    }

    updateValue(newValue: any) {
        console.log("Updated", this);
        this.onChange.next(true);
    }

    toUTC(value: number): number {
        let date = new Date();
        return (value + date.getTimezoneOffset() + 10080) % 10080;
    }

    fromUTC(value: number): number {
        let date = new Date();
        return (value - date.getTimezoneOffset() + 10080) % 10080;
    }

    closestMultiple(value: number, multiple: number) : number {
        return ((value / multiple) | 0) * multiple;
    }

    toIntervals(): Interval[] {
        return this.days.map((value) => {
            let start = this.toUTC(value * 1440 + this.start);
            let end = this.toUTC(value * 1440 + this.end);
            console.log(value, start, this.start, end, this.end);
            let duration = end - start;
            return new Interval(start, duration);
        })
    }

    static toIntervals(interfaces: IntervalInterface[]): Interval[] {
        return interfaces.reduce((acc: Interval[], val: IntervalInterface) => {
            val.toIntervals().filter( (val) => {
                return acc.findIndex( (elem) => elem.duration == val.duration && elem.start == val.start) < 0
            }).forEach( elem => {
                acc.push(elem);
            })

            return acc;
        }, [])
    }

    static convertIntervals(interfaces: IntervalInterface[], intervals: Interval[]) {
        interfaces.length = 0;
        intervals.reduce((prev: IntervalInterface[], next: Interval) => {
            let interval = new IntervalInterface(next);
            let duplicate = prev.find((value: IntervalInterface) => {
                return value.start === interval.start && value.end === interval.end;
            })
            if (duplicate) {
                if ( duplicate.days.find( value => value == interval.days[0] ) === undefined) {
                    duplicate.days.push(interval.days[0]);
                }
            } else {
                prev.push(interval);
            }
            return prev;
        }, []).forEach((value) => {
            interfaces.push(value);
        })
    }
}
