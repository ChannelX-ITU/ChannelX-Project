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
    MatIconModule
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
      MatGridListModule,
      MatIconModule,
      NoopAnimationsModule
  ],
  declarations: []
})
export class MaterialModulesModule { }
