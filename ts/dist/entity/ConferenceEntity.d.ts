import { NhlApiDocumentationEntityBase } from '../NhlApiDocumentationEntityBase';
import type { NhlApiDocumentationSDK } from '../NhlApiDocumentationSDK';
import type { Control } from '../types';
import type { Conference, ConferenceLoadMatch, ConferenceListMatch } from '../NhlApiDocumentationTypes';
declare class ConferenceEntity extends NhlApiDocumentationEntityBase<Conference> {
    constructor(client: NhlApiDocumentationSDK, entopts: any);
    make(this: ConferenceEntity): ConferenceEntity;
    load(this: any, reqmatch?: ConferenceLoadMatch, ctrl?: Control): Promise<ConferenceEntity>;
    list(this: any, reqmatch?: ConferenceListMatch, ctrl?: Control): Promise<ConferenceEntity[]>;
}
export { ConferenceEntity };
