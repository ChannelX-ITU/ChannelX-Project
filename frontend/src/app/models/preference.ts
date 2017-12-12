import { Interval } from './interval'

export class Preference {
    constructor(
        public start_date: number = Date.now(),
        public duration: number = 0,
        public intervals: Interval[] = [])
    {}
}
