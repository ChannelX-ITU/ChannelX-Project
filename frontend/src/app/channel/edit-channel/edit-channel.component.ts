import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { IndexerPipe } from '../../pipes/indexer.pipe';
import { IntervalInterface } from '../../interval/interval.component'
import { Logger } from '@nsalaun/ng-logger'
@Component({
  selector: 'app-edit-channel',
  templateUrl: './edit-channel.component.html',
  styleUrls: ['./edit-channel.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class EditChannelComponent implements OnInit {

  intervals: IntervalInterface[] = [
    new IntervalInterface()
  ];

  constructor(private logger: Logger) { }

  ngOnInit() {
  }

  createNew()
  {
    this.intervals.push(new IntervalInterface())
    this.logger.log(this.intervals);
  }
}
