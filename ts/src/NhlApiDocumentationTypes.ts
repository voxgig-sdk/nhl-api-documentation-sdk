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

export interface ConferenceListMatch {
  conference?: any[]
  copyright?: string
  id?: number
  link?: string
  name?: string
}

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

export interface DivisionListMatch {
  copyright?: string
  division?: any[]
  id?: number
  link?: string
  name?: string
}

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

export interface ScheduleListMatch {
  date?: string
  game?: any[]
  total_event?: number
  total_game?: number
  total_item?: number
  total_match?: number
}

export interface Standing {
  conference?: Record<string, any>
  division?: Record<string, any>
  team_record?: any[]
}

export interface StandingListMatch {
  conference?: Record<string, any>
  division?: Record<string, any>
  team_record?: any[]
}

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

export interface TeamListMatch {
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

