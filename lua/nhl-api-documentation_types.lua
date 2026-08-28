-- Typed models for the NhlApiDocumentation SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Conference
---@field conferences? table
---@field copyright? string
---@field id? number
---@field link? string
---@field name? string

---@class ConferenceLoadMatch
---@field id number

---@class ConferenceListMatch
---@field conferences? table
---@field copyright? string
---@field id? number
---@field link? string
---@field name? string

---@class Division
---@field copyright? string
---@field divisions? table
---@field id? number
---@field link? string
---@field name? string

---@class DivisionLoadMatch
---@field id number

---@class DivisionListMatch
---@field copyright? string
---@field divisions? table
---@field id? number
---@field link? string
---@field name? string

---@class Game
---@field away? table
---@field copyright? string
---@field gameData? table
---@field gamePk? number
---@field home? table
---@field id? string
---@field link? string
---@field liveData? table

---@class GameLoadMatch
---@field id number

---@class Player
---@field copyright? string
---@field id? string
---@field people? table

---@class PlayerLoadMatch
---@field id number

---@class PlayerStat
---@field splits? table
---@field type? table

---@class PlayerStatListMatch
---@field person_id number
---@field season? string
---@field stat string

---@class Roster
---@field jerseyNumber? string
---@field person? table
---@field position? table

---@class RosterListMatch
---@field team_id number
---@field season? string

---@class Schedule
---@field date? string
---@field games? table
---@field totalEvents? number
---@field totalGames? number
---@field totalItems? number
---@field totalMatches? number

---@class ScheduleListMatch
---@field end_date? string
---@field season? string
---@field start_date? string
---@field team_id? number

---@class Standing
---@field conference? table
---@field division? table
---@field teamRecords? table

---@class StandingListMatch
---@field season? string

---@class Team
---@field abbreviation? string
---@field conference? table
---@field copyright? string
---@field division? table
---@field firstYearOfPlay? string
---@field franchise? table
---@field id? number
---@field link? string
---@field locationName? string
---@field name? string
---@field teamName? string
---@field teams? table
---@field venue? table

---@class TeamLoadMatch
---@field id number
---@field expand? string

---@class TeamListMatch
---@field expand? string
---@field season? string

local M = {}

return M
