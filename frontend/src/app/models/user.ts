import { Channel } from './channel'
import { Preference } from './preference'
import { Communication } from './communication'

export class User {

    constructor(
        public username: string = "", 
        public channels: Channel[] = [], 
        public preferences: Preference = new Preference(), 
        public communications: Communication[] = [])
    {}

}
