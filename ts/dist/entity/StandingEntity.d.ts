import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Standing, StandingListMatch } from '../NhlApiDocumentationTypes';
declare class StandingEntity extends NhlApiDocumentationEntityBase<Standing> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: StandingEntity): StandingEntity;
    list(this: any, reqmatch?: StandingListMatch, ctrl?: Control): Promise<StandingEntity[]>;
}
export { StandingEntity };
