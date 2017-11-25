import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class MessageComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
