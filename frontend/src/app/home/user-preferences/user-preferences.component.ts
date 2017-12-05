import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';

@Component({
  selector: 'app-user-preferences',
  templateUrl: './user-preferences.component.html',
  styleUrls: ['./user-preferences.component.scss']
})
export class UserPreferencesComponent implements OnInit {

    emails = [
      {value: '0', name: 'Default Mail', view: 'ihsanyigitergin@gmail.com'},
      {value: '1', name: 'School Mail', view: 'erginihs@itu.edu.tr'}
    ];

    phones = [
      {value: '0', name: 'Default Phone', view: '+905546544365'},
      {value: '1', name: 'School Phone', view: '+905318893565'}
    ];

  constructor() { }

  ngOnInit() {
  }

}
