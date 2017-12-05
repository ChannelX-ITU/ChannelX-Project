import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { HomeComponent } from './home/home.component';
import { ChannelComponent } from './channel/channel.component';
import { ChannelsComponent } from './home/channels/channels.component';
import { CreateChannelComponent } from './home/create-channel/create-channel.component';
import { UserPreferencesComponent } from './home/user-preferences/user-preferences.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent} from './auth/register/register.component'
import { AuthComponent } from './auth/auth.component'
import { AuthGuard } from './guards/auth.guard'

const routes: Routes = [
  { path: '', component: AuthComponent, children: [
      { path: '', redirectTo: 'login', pathMatch: 'full' },
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent }
    ]
  },
  { path: 'about', component: AboutComponent },
  { path: 'home', component: HomeComponent,
      canActivate: [ AuthGuard ],
      children: [
          { path: "", redirectTo: "channels", pathMatch: "full" },
          { path: "channels", component: ChannelsComponent },
          { path: "preferences", component: UserPreferencesComponent },
          { path: "create", component: CreateChannelComponent }
      ]
  },
  { path: 'channel', component: ChannelComponent },
  { path: '**', redirectTo: '' }
];


@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
