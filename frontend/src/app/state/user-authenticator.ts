import { Action } from '@ngrx/store';
import { UserState, LoginAction } from './user-state'
import { User } from '../models/user';

const LOGIN = 'LOGIN'
const LOGOUT = 'LOGOUT'

export function userAuth(state: UserState = new UserState(), action: LoginAction) {
    switch (action.type) {
        case LOGIN:
            return new UserState(action.user, true);
        case LOGOUT:
            return new UserState(null, false);
        default:
            return new UserState(null, false);
    }
}
