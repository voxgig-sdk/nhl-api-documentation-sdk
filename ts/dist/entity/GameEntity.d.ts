import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Game, GameLoadMatch } from '../NhlApiDocumentationTypes';
declare class GameEntity extends NhlApiDocumentationEntityBase<Game> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: GameEntity): GameEntity;
    load(this: any, reqmatch?: GameLoadMatch, ctrl?: Control): Promise<GameEntity>;
}
export { GameEntity };
