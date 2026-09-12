import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Division, DivisionLoadMatch, DivisionListMatch } from '../NhlApiDocumentationTypes';
declare class DivisionEntity extends NhlApiDocumentationEntityBase<Division> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: DivisionEntity): DivisionEntity;
    load(this: any, reqmatch?: DivisionLoadMatch, ctrl?: Control): Promise<DivisionEntity>;
    list(this: any, reqmatch?: DivisionListMatch, ctrl?: Control): Promise<DivisionEntity[]>;
}
export { DivisionEntity };
