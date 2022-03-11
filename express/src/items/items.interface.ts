// plural items

import {Item} from './item.interface';

// store a bunch of item objects in items
export interface Items {
    [key: number]: Item;
}