import { Channel } from './channel'
import { Preference } from './preference'
import { Communication } from './communication'

export class User {
    username: string;
    channels: Channel[];
    preferences: Preference;
    communications: Communication[];
}
