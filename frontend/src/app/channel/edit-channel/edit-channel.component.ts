import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { IndexerPipe } from '../../pipes/indexer.pipe';

@Component({
  selector: 'app-edit-channel',
  templateUrl: './edit-channel.component.html',
  styleUrls: ['./edit-channel.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class EditChannelComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  createNew()
  {
    console.log("Whoa");
  }
}
