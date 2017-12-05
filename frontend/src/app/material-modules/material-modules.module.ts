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
  MatTableModule,
  MatExpansionModule,
  MatTooltipModule
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
    MatTableModule,
    MatExpansionModule,
    MatTooltipModule
  ],
  declarations: []
})
export class MaterialModulesModule { }
