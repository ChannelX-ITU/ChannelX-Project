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
  MatTooltipModule,
  MatProgressSpinnerModule,
  MatCheckboxModule
} from '@angular/material';

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
	  MatOptionModule,
	  MatSelectModule,
	  MatNativeDateModule,
	  MatDatepickerModule,
	  MatSliderModule,
	  MatStepperModule,
    MatTableModule,
    MatExpansionModule,
    MatTooltipModule,
    MatProgressSpinnerModule,
    MatCheckboxModule
  ],
  declarations: []
})
export class MaterialModulesModule { }
