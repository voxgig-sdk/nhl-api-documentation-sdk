// Typed models for the NhlApiDocumentation SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Conference {
  conference?: any[]
  copyright?: string
  id?: number
  link?: string
  name?: string
}

export interface ConferenceLoadMatch {
  id: number
}

export type ConferenceListMatch = Partial<Conference>

export interface Division {
  copyright?: string
  division?: any[]
  id?: number
  link?: string
  name?: string
}

export interface DivisionLoadMatch {
  id: number
}

export type DivisionListMatch = Partial<Division>

export interface Game {
  copyright?: string
  game_data?: Record<string, any>
  game_pk?: number
  link?: string
  live_data?: Record<string, any>
  team?: Record<string, any>
}

export interface GameLoadMatch {
  id: number
}

export interface Player {
  copyright?: string
  person?: any[]
}

export interface PlayerLoadMatch {
  id: number
}

export interface PlayerStat {
  split?: any[]
  type?: Record<string, any>
}

export interface PlayerStatListMatch {
  person_id: number
}

export interface Roster {
  jersey_number?: string
  person?: Record<string, any>
  position?: Record<string, any>
}

export interface RosterListMatch {
  team_id: number
}

export interface Schedule {
  date?: string
  game?: any[]
  total_event?: number
  total_game?: number
  total_item?: number
  total_match?: number
}

export type ScheduleListMatch = Partial<Schedule>

export interface Standing {
  conference?: Record<string, any>
  division?: Record<string, any>
  team_record?: any[]
}

export type StandingListMatch = Partial<Standing>

export interface Team {
  abbreviation?: string
  conference?: Record<string, any>
  copyright?: string
  division?: Record<string, any>
  first_year_of_play?: string
  franchise?: Record<string, any>
  id?: number
  link?: string
  location_name?: string
  name?: string
  team?: any[]
  team_name?: string
  venue?: Record<string, any>
}

export interface TeamLoadMatch {
  id: number
}

export type TeamListMatch = Partial<Team>

