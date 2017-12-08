import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import { IndexerPipe } from '../../pipes/indexer.pipe';
import { IntervalInterface } from '../../interval/interval.component'
import { Interval } from '../../models/interval';
import { Logger } from '@nsalaun/ng-logger'
import { Channel } from '../../models/channel';
import { Observable } from 'rxjs/Observable';
import { AppState } from '../../state/app-state';
import { Store } from '@ngrx/store';
import { map } from 'rxjs/operators';
import { Communication } from '../../models/communication'
import { RouteChildBinderService } from '../../services/route-child-binder.service'


@Component({
  selector: 'app-edit-channel',
  templateUrl: './edit-channel.component.html',
  styleUrls: ['./edit-channel.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class EditChannelComponent implements OnInit {

  channel: Observable<Channel>;

  comms: Observable<Communication[]>;
  intervals: IntervalInterface[] = [];

  constructor(
    private logger: Logger, 
    private store: Store<AppState>, 
    private childBinder: RouteChildBinderService<Channel, boolean>) { }

  ngOnInit() {
    this.childBinder.fromParent.subscribe(value => {
      this.logger.log("child:", value);
      this.convertIntervals(value)
    });
    this.comms = this.store
    .select("user")
    .pipe(map( 
      (value) => value.user ? value.user.communications : new Array<Communication>()
      )
    );
  }

  save() {
    let intervals = this.intervals.map( value => value.toIntervals()).reduce((acc, value) => {
      value.forEach( value => acc.push(value));
      return acc;
    }, []);
    this.logger.log("Intervals: " , intervals);
  }

  convertIntervals(channel: Channel) {
    this.logger.log("convert:", channel);
    if (!channel) return;
    this.intervals.length = 0;
    channel.preference.intervals.reduce((prev: IntervalInterface[], next: Interval) => {
      let interval = new IntervalInterface(next);
      this.logger.log("Created ", interval, "from", next);
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
      this.logger.log("interval: ", value);
      this.intervals.push(value);
    })
  }

  createNew()
  {
    this.intervals.push(new IntervalInterface())
    this.logger.log(this.intervals);
  }
}
