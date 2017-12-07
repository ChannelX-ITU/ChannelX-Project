import { Preference } from './preference';
import { Restriction } from './restriction';

export class Channel {
    name: string;
    is_owner: boolean;
    preference: Preference;
    users: string[];
    restrictions: Restriction[];
}
