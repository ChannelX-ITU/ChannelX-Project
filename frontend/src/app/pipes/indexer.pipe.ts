import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'indexer'
})
export class IndexerPipe implements PipeTransform {

  transform(value: number, args?: any): number[] {
    let arr = Array(value).fill(1);
    for(let i = 0; i < arr.length; ++i) {
        arr[i] = i;
    }
    return arr;
  }

}
