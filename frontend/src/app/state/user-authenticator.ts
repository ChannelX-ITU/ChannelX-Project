import { Action } from '@ngrx/store';
import { UserState, LoginAction } from './user-state'
import { User } from '../models/user';

const LOGIN = 'LOGIN'
const LOGOUT = 'LOGOUT'
const AUTO_LOGIN = 'AUTO_LOGIN'
const REFRESH = 'REFRESH'

export function userAuth(state: UserState = new UserState(), action: LoginAction) {
    switch (action.type) {
        case LOGIN:
            return new UserState(action.user, true);
        case AUTO_LOGIN:
            return new UserState(action.user, true);
        case REFRESH:
            return new UserState(action.user, true);
        case LOGOUT:
            return new UserState(null, false);
        default:
            return new UserState(null, false);
    }
}
