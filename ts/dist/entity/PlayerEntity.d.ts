import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Player, PlayerLoadMatch } from '../NhlApiDocumentationTypes';
declare class PlayerEntity extends NhlApiDocumentationEntityBase<Player> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: PlayerEntity): PlayerEntity;
    load(this: any, reqmatch?: PlayerLoadMatch, ctrl?: Control): Promise<PlayerEntity>;
}
export { PlayerEntity };
