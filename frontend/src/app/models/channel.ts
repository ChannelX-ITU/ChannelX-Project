import { Preference } from './preference';
import { Restriction } from './restriction';
import { Communication} from './communication';

export class Channel {
    constructor(
        public name: string = "",
        public is_owner: boolean = false,
        public preference: Preference = new Preference(),
        public users: string[] = [],
        public restrictions: Restriction[] = [],
        public comm: string = ""
        ) {}
}

export class ChannelResponse {
    constructor(
        public alias: string = "",
        public channel: Channel = new Channel(),
        public comm: Communication = new Communication()
        ) {
        this.channel.comm = comm.value;
    }
}
