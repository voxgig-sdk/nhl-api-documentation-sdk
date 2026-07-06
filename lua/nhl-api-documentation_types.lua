-- Typed models for the NhlApiDocumentation SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Conference
---@field conference? table
---@field copyright? string
---@field id? number
---@field link? string
---@field name? string

---@class ConferenceLoadMatch
---@field id number

---@class ConferenceListMatch
---@field conference? table
---@field copyright? string
---@field id? number
---@field link? string
---@field name? string

---@class Division
---@field copyright? string
---@field division? table
---@field id? number
---@field link? string
---@field name? string

---@class DivisionLoadMatch
---@field id number

---@class DivisionListMatch
---@field copyright? string
---@field division? table
---@field id? number
---@field link? string
---@field name? string

---@class Game
---@field copyright? string
---@field game_data? table
---@field game_pk? number
---@field link? string
---@field live_data? table
---@field team? table

---@class GameLoadMatch
---@field id number

---@class Player
---@field copyright? string
---@field person? table

---@class PlayerLoadMatch
---@field id number

---@class PlayerStat
---@field split? table
---@field type? table

---@class PlayerStatListMatch
---@field person_id number

---@class Roster
---@field jersey_number? string
---@field person? table
---@field position? table

---@class RosterListMatch
---@field team_id number

---@class Schedule
---@field date? string
---@field game? table
---@field total_event? number
---@field total_game? number
---@field total_item? number
---@field total_match? number

---@class ScheduleListMatch
---@field date? string
---@field game? table
---@field total_event? number
---@field total_game? number
---@field total_item? number
---@field total_match? number

---@class Standing
---@field conference? table
---@field division? table
---@field team_record? table

---@class StandingListMatch
---@field conference? table
---@field division? table
---@field team_record? table

---@class Team
---@field abbreviation? string
---@field conference? table
---@field copyright? string
---@field division? table
---@field first_year_of_play? string
---@field franchise? table
---@field id? number
---@field link? string
---@field location_name? string
---@field name? string
---@field team? table
---@field team_name? string
---@field venue? table

---@class TeamLoadMatch
---@field id number

---@class TeamListMatch
---@field abbreviation? string
---@field conference? table
---@field copyright? string
---@field division? table
---@field first_year_of_play? string
---@field franchise? table
---@field id? number
---@field link? string
---@field location_name? string
---@field name? string
---@field team? table
---@field team_name? string
---@field venue? table

local M = {}

return M
