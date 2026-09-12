import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Roster, RosterListMatch } from '../NhlApiDocumentationTypes';
declare class RosterEntity extends NhlApiDocumentationEntityBase<Roster> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: RosterEntity): RosterEntity;
    list(this: any, reqmatch?: RosterListMatch, ctrl?: Control): Promise<RosterEntity[]>;
}
export { RosterEntity };
