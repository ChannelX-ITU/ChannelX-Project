import { Pipe, PipeTransform } from '@angular/core';
import { Communication } from '../models/communication';
import { Logger} from '@nsalaun/ng-logger';

@Pipe({
  name: 'commFilter'
})
export class CommFilterPipe implements PipeTransform {

    constructor(private logger: Logger) {}

  transform(value: Communication[], type: String): Communication[] {
      this.logger.log("Filtering", value, "with", type);
    return value.filter( val => val.comm_type === type);
  }

}
