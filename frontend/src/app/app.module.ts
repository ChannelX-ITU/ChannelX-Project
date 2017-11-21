import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { MaterialModulesModule } from './material-modules/material-modules.module';
import { AboutComponent } from './about/about.component';
import { AppRoutingModule } from './/app-routing.module';
import { HomeComponent } from './home/home.component'

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    ToolbarComponent,
    AboutComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    MaterialModulesModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
