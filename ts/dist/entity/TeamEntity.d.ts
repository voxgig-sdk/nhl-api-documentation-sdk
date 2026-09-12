import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Team, TeamLoadMatch, TeamListMatch } from '../NhlApiDocumentationTypes';
declare class TeamEntity extends NhlApiDocumentationEntityBase<Team> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: TeamEntity): TeamEntity;
    load(this: any, reqmatch?: TeamLoadMatch, ctrl?: Control): Promise<TeamEntity>;
    list(this: any, reqmatch?: TeamListMatch, ctrl?: Control): Promise<TeamEntity[]>;
}
export { TeamEntity };
