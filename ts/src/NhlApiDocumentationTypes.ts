// Typed models for the NhlApiDocumentation SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Conference {
  conferences?: any[]
  copyright?: string
  id?: number
  link?: string
  name?: string
}

export interface ConferenceLoadMatch {
  id: number
}

export interface ConferenceListMatch {
  conferences?: any[]
  copyright?: string
  id?: number
  link?: string
  name?: string
}

export interface Division {
  copyright?: string
  divisions?: any[]
  id?: number
  link?: string
  name?: string
}

export interface DivisionLoadMatch {
  id: number
}

export interface DivisionListMatch {
  copyright?: string
  divisions?: any[]
  id?: number
  link?: string
  name?: string
}

export interface Game {
  away?: Record<string, any>
  copyright?: string
  gameData?: Record<string, any>
  gamePk?: number
  home?: Record<string, any>
  id?: string
  link?: string
  liveData?: Record<string, any>
}

export interface GameLoadMatch {
  id: number

  // Selects a custom action instead of the plain load:
  //   'boxscore' | 'feed_live'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Player {
  copyright?: string
  id?: string
  people?: any[]
}

export interface PlayerLoadMatch {
  id: number
}

export interface PlayerStat {
  splits?: any[]
  type?: Record<string, any>
}

export interface PlayerStatListMatch {
  person_id: number
  season?: string
  stat: string
}

export interface Roster {
  jerseyNumber?: string
  person?: Record<string, any>
  position?: Record<string, any>
}

export interface RosterListMatch {
  team_id: number
  season?: string
}

export interface Schedule {
  date?: string
  games?: any[]
  totalEvents?: number
  totalGames?: number
  totalItems?: number
  totalMatches?: number
}

export interface ScheduleListMatch {
  end_date?: string
  season?: string
  start_date?: string
  team_id?: number
}

export interface Standing {
  conference?: Record<string, any>
  division?: Record<string, any>
  teamRecords?: any[]
}

export interface StandingListMatch {
  season?: string
}

export interface Team {
  abbreviation?: string
  conference?: Record<string, any>
  copyright?: string
  division?: Record<string, any>
  firstYearOfPlay?: string
  franchise?: Record<string, any>
  id?: number
  link?: string
  locationName?: string
  name?: string
  teamName?: string
  teams?: any[]
  venue?: Record<string, any>
}

export interface TeamLoadMatch {
  id: number
  expand?: string
}

export interface TeamListMatch {
  expand?: string
  season?: string
}

