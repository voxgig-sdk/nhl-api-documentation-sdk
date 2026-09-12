import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Schedule, ScheduleListMatch } from '../NhlApiDocumentationTypes';
declare class ScheduleEntity extends NhlApiDocumentationEntityBase<Schedule> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: ScheduleEntity): ScheduleEntity;
    list(this: any, reqmatch?: ScheduleListMatch, ctrl?: Control): Promise<ScheduleEntity[]>;
}
export { ScheduleEntity };
