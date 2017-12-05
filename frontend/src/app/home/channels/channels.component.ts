import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import {MatTableDataSource} from '@angular/material';


@Component({
  selector: 'app-channels',
  templateUrl: './channels.component.html',
  styleUrls: ['./channels.component.scss']
})
export class ChannelsComponent implements OnInit {

    displayedColumns = ['ChannelName', 'UserCount', 'CommType'];
    dataSource = new MatTableDataSource<Element>(ELEMENT_DATA);

  constructor() { }

  ngOnInit() {
  }

}

export interface Element {
  ChannelName: string;
  UserCount: number;
  CommType: string;
  IsActive: number;
}

const ELEMENT_DATA: Element[] = [
  {ChannelName: 'BluePanda', UserCount: 10, CommType: 'Email', IsActive: 1},
  {ChannelName: 'RedIguana', UserCount: 12, CommType: 'Email', IsActive: 1},
  {ChannelName: 'GreyChipmunk', UserCount: 23, CommType: 'Email', IsActive: 1},
  {ChannelName: 'FlyingWhale', UserCount: 2, CommType: 'SMS', IsActive: 1}
];
