import {ContentChild, Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-interval',
  templateUrl: './interval.component.html',
  styleUrls: ['./interval.component.scss'],
  encapsulation: ViewEncapsulation.None,
  template: `
    <template #temp>
    <mat-card id="interval">
      <mat-form-field>
        <mat-select placeholder="Choose Active Days" multiple>
          <mat-option *ngFor="let day of days" [value]="day.view">
          {{day.view}}
          </mat-option>
        </mat-select>
      </mat-form-field>
      <br/>
      <mat-form-field>
        <mat-select placeholder="Choose Active Hour Begin">
          <mat-option *ngFor="let hour of hours" [value]="hour.view">
          {{hour.view}}
          </mat-option>
        </mat-select>
      </mat-form-field>
      <br/>
      <mat-form-field>
        <mat-select placeholder="Choose Active Hour Finish">
          <mat-option *ngFor="let hour of hours" [value]="hour.view">
          {{hour.view}}
          </mat-option>
        </mat-select>
      </mat-form-field>
      <button mat-fab (click)="createNew()">+</button>
    </mat-card>
    </template>
    <template [ngTemplateOutlet]="temp"></template>
  `
})
export class IntervalComponent implements OnInit {
  bla: boolean = false;
      @ContentChild('temp') testEl: any;
  constructor() { }

  ngOnInit() {
  }

}
