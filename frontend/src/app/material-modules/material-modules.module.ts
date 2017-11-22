import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
    MatCardModule,
    MatTabsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatToolbarModule,
    MatGridListModule,
    MatIconModule,
	MatOptionModule,
	MatSelectModule,
	MatNativeDateModule,
	MatDatepickerModule,
	MatSliderModule,
  MatListModule,
	MatStepperModule,
  MatTableModule
} from '@angular/material';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [
      MatCardModule,
      MatTabsModule,
      MatFormFieldModule,
      MatInputModule,
      MatButtonModule,
      MatToolbarModule,
      MatListModule,
      MatGridListModule,
      MatIconModule,
      NoopAnimationsModule,
	  MatOptionModule,
	  MatSelectModule,
	  MatNativeDateModule,
	  MatDatepickerModule,
	  MatSliderModule,
	  MatStepperModule,
    MatTableModule
  ],
  declarations: []
})
export class MaterialModulesModule { }
