import { Action } from '@ngrx/store';
import { UserState } from './user-state'

const LOGIN = 'LOGIN'
const LOGOUT = 'LOGOUT'

export function userAuth(state: UserState = new UserState(), action: Action) {
    switch (action.type) {
        case LOGIN:
            return new UserState(state.user, true);
        case LOGOUT:
            return new UserState(state.user, false);
        default:
            return new UserState(null, false);
    }
}
