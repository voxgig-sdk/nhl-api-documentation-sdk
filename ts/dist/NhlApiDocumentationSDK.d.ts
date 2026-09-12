import { ConferenceEntity } from './entity/ConferenceEntity';
import { DivisionEntity } from './entity/DivisionEntity';
import { GameEntity } from './entity/GameEntity';
import { PlayerEntity } from './entity/PlayerEntity';
import { PlayerStatEntity } from './entity/PlayerStatEntity';
import { RosterEntity } from './entity/RosterEntity';
import { ScheduleEntity } from './entity/ScheduleEntity';
import { StandingEntity } from './entity/StandingEntity';
import { TeamEntity } from './entity/TeamEntity';
export type * from './NhlApiDocumentationTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { NhlApiDocumentationEntityBase } from './NhlApiDocumentationEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class NhlApiDocumentationSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Conference(entopts?: Record<string, any>): ConferenceEntity;
    Division(entopts?: Record<string, any>): DivisionEntity;
    Game(entopts?: Record<string, any>): GameEntity;
    Player(entopts?: Record<string, any>): PlayerEntity;
    PlayerStat(entopts?: Record<string, any>): PlayerStatEntity;
    Roster(entopts?: Record<string, any>): RosterEntity;
    Schedule(entopts?: Record<string, any>): ScheduleEntity;
    Standing(entopts?: Record<string, any>): StandingEntity;
    Team(entopts?: Record<string, any>): TeamEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): NhlApiDocumentationSDK;
    tester(testopts?: any, sdkopts?: any): NhlApiDocumentationSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof NhlApiDocumentationSDK;
export { stdutil, config, BaseFeature, NhlApiDocumentationEntityBase, NhlApiDocumentationSDK, SDK, };
