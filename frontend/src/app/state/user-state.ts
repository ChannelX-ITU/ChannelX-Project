import { User } from '../models/user'

export class UserState {
    constructor(public user: User = null, public logged_in: boolean = false) {
        
    }
}

export interface LoginAction {
    type: string,
    user: User
} 
