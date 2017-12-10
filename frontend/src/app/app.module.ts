import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { SimpleNotificationsModule } from 'angular2-notifications';


import { AppComponent } from './app.component';
import { LoginComponent } from './auth/login/login.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { MaterialModulesModule } from './material-modules/material-modules.module';
import { AboutComponent } from './about/about.component';
import { AppRoutingModule } from './/app-routing.module';
import { HomeComponent } from './home/home.component';
import { ChannelComponent } from './channel/channel.component';
import { ReplyComponent } from './reply/reply.component';
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
import { ActivationGuard } from './guards/activation.guard';
import { BroadcastComponent } from './channel/broadcast/broadcast.component';
import { EditChannelComponent } from './channel/edit-channel/edit-channel.component';
import { IndexerPipe } from './pipes/indexer.pipe';
import { IntervalComponent } from './interval/interval.component';
import { CommFilterPipe } from './pipes/comm-filter.pipe';
import { RouteChildBinderService } from './services/route-child-binder.service';
import { FloorPipe } from './pipes/floor.pipe';
import { CeilPipe } from './pipes/ceil.pipe'

import { ErrorNotifyInterceptor } from './interceptors/error-notify.interceptor'
import { SuccessNotifyInterceptor } from './interceptors/success-notify.interceptor';
import { PreferencesComponent } from './preferences/preferences.component';
import { CommunicationsComponent } from './communications/communications.component';

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
    ReplyComponent,
    ChannelComponent,
    UserPreferencesComponent,
    ChannelsComponent,
    CreateChannelComponent,
    AuthComponent,
    RegisterComponent,
    BroadcastComponent,
    EditChannelComponent,
    IndexerPipe,
    IntervalComponent,
    CommFilterPipe,
    FloorPipe,
    CeilPipe,
    PreferencesComponent,
    CommunicationsComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MaterialModulesModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    NgLoggerModule.forRoot(LOG_LEVEL),
    CookieModule.forRoot(),
    StoreModule.forRoot({ user: userAuth }),
    SimpleNotificationsModule.forRoot()
  ],
  providers: [ AuthGuard, ActivationGuard, RouteChildBinderService, {
    provide: HTTP_INTERCEPTORS,
    useClass: ErrorNotifyInterceptor,
    multi: true,
  },{
    provide: HTTP_INTERCEPTORS,
    useClass: SuccessNotifyInterceptor,
    multi: true,
  } ],
  bootstrap: [AppComponent]
})
export class AppModule { }
