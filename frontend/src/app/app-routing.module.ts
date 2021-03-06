import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutComponent } from './about/about.component';
import { HomeComponent } from './home/home.component';
import { ChannelsComponent } from './home/channels/channels.component';
import { CreateChannelComponent } from './home/create-channel/create-channel.component';
import { UserPreferencesComponent } from './home/user-preferences/user-preferences.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent} from './auth/register/register.component'
import { AuthComponent } from './auth/auth.component'

import { ChannelComponent } from './channel/channel.component'
import { UsersComponent } from './channel/users/users.component';

import { BroadcastComponent } from './channel/broadcast/broadcast.component'
import { EditChannelComponent } from './channel/edit-channel/edit-channel.component'
import { ReplyComponent } from './reply/reply.component'

import { AuthGuard } from './guards/auth.guard'
import { ActivationGuard } from './guards/activation.guard'

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
  { path: 'channel/:id', component: ChannelComponent,
      canActivate: [ AuthGuard ],
      children: [
        { path: "", redirectTo: "broadcast", pathMatch: "full" },
        { path: "settings", component: EditChannelComponent },
        { path: "broadcast", component: BroadcastComponent },
        { path: "users", component: UsersComponent }
      ] 
  },
  {
    path: 'activate/:token', canActivate: [ ActivationGuard ], component: AuthComponent
  },
  {
    path: 'reply', component: ReplyComponent
  },
  { path: '**', redirectTo: '' }
];


@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
