import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { AboutComponent } from './about/about.component';
import { HomeComponent } from './home/home.component';
import { ChannelComponent } from './channel/channel.component';
import { ChannelsComponent } from './home/channels/channels.component';
import { CreateChannelComponent } from './home/create-channel/create-channel.component';
import { UserPreferencesComponent } from './home/user-preferences/user-preferences.component';

const routes: Routes = [
  { path: '', redirectTo: 'about', pathMatch: 'full'},
  { path: 'login', component: LoginComponent },
  { path: 'about', component: AboutComponent },
  { path: 'home', component: HomeComponent,
      children: [
          { path: "", redirectTo: "channels", pathMatch: "full" },
          { path: "channels", component: ChannelsComponent },
          { path: "preferences", component: UserPreferencesComponent },
          { path: "create", component: CreateChannelComponent }
      ]
  },
  { path: 'channel', component: ChannelComponent }
];


@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
