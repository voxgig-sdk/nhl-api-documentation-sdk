import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { PlayerStat, PlayerStatListMatch } from '../NhlApiDocumentationTypes';
declare class PlayerStatEntity extends NhlApiDocumentationEntityBase<PlayerStat> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: PlayerStatEntity): PlayerStatEntity;
    list(this: any, reqmatch?: PlayerStatListMatch, ctrl?: Control): Promise<PlayerStatEntity[]>;
}
export { PlayerStatEntity };
