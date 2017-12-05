import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { SimpleNotificationsModule } from 'angular2-notifications';


import { AppComponent } from './app.component';
import { LoginComponent } from './auth/login/login.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { MaterialModulesModule } from './material-modules/material-modules.module';
import { AboutComponent } from './about/about.component';
import { AppRoutingModule } from './/app-routing.module';
import { HomeComponent } from './home/home.component';
import { ChannelComponent } from './channel/channel.component';
import { UserPreferencesComponent } from './home/user-preferences/user-preferences.component';
import { ChannelsComponent } from './home/channels/channels.component';
import { CreateChannelComponent } from './home/create-channel/create-channel.component';
import { AuthComponent } from './auth/auth.component';
import { RegisterComponent } from './auth/register/register.component';
import { NgLoggerModule, Level } from '@nsalaun/ng-logger';
import { CookieModule } from 'ngx-cookie';
import { environment } from '../environments/environment';

import { StoreModule } from '@ngrx/store';
import { userAuth } from './state/user-authenticator'
import { AuthGuard } from './guards/auth.guard';
import { BroadcastComponent } from './channel/broadcast/broadcast.component';
import { EditChannelComponent } from './channel/edit-channel/edit-channel.component';
import { IndexerPipe } from './pipes/indexer.pipe'

const LOG_LEVEL = Level.LOG;
if (environment.production){
    const LOG_LEVEL = Level.ERROR;
}

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    ToolbarComponent,
    AboutComponent,
    HomeComponent,
    ChannelComponent,
    UserPreferencesComponent,
    ChannelsComponent,
    CreateChannelComponent,
    AuthComponent,
    RegisterComponent,
    BroadcastComponent,
    EditChannelComponent,
    IndexerPipe
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MaterialModulesModule,
    AppRoutingModule,
    FormsModule,
    NgLoggerModule.forRoot(LOG_LEVEL),
    CookieModule.forRoot(),
    StoreModule.forRoot({ user: userAuth }),
    SimpleNotificationsModule.forRoot()
  ],
  providers: [ AuthGuard ],
  bootstrap: [AppComponent]
})
export class AppModule { }
