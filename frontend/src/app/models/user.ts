import { Channel } from './channel'
import { Preference } from './preference'
import { CommMethod } from './comm-method'

export class User {
    username: string;
    channels: Channel[];
    preferences: Preference;
    comm_methods: CommMethod;
}
