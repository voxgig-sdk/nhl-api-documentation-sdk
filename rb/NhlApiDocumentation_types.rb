# frozen_string_literal: true

# Typed models for the NhlApiDocumentation SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Conference entity data model.
#
# @!attribute [rw] conferences
#   @return [Array, nil]
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
Conference = Struct.new(
  :conferences,
  :copyright,
  :id,
  :link,
  :name,
  keyword_init: true
)

# Request payload for Conference#load.
#
# @!attribute [rw] id
#   @return [Integer]
ConferenceLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Conference#list.
#
# @!attribute [rw] conferences
#   @return [Array, nil]
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
ConferenceListMatch = Struct.new(
  :conferences,
  :copyright,
  :id,
  :link,
  :name,
  keyword_init: true
)

# Division entity data model.
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] divisions
#   @return [Array, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
Division = Struct.new(
  :copyright,
  :divisions,
  :id,
  :link,
  :name,
  keyword_init: true
)

# Request payload for Division#load.
#
# @!attribute [rw] id
#   @return [Integer]
DivisionLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Division#list.
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] divisions
#   @return [Array, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
DivisionListMatch = Struct.new(
  :copyright,
  :divisions,
  :id,
  :link,
  :name,
  keyword_init: true
)

# Game entity data model.
#
# @!attribute [rw] away
#   @return [Hash, nil]
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] gameData
#   @return [Hash, nil]
#
# @!attribute [rw] gamePk
#   @return [Integer, nil]
#
# @!attribute [rw] home
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] liveData
#   @return [Hash, nil]
Game = Struct.new(
  :away,
  :copyright,
  :gameData,
  :gamePk,
  :home,
  :id,
  :link,
  :liveData,
  keyword_init: true
)

# Request payload for Game#load.
#
# @!attribute [rw] id
#   @return [Integer]
GameLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Player entity data model.
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] people
#   @return [Array, nil]
Player = Struct.new(
  :copyright,
  :id,
  :people,
  keyword_init: true
)

# Request payload for Player#load.
#
# @!attribute [rw] id
#   @return [Integer]
PlayerLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# PlayerStat entity data model.
#
# @!attribute [rw] splits
#   @return [Array, nil]
#
# @!attribute [rw] type
#   @return [Hash, nil]
PlayerStat = Struct.new(
  :splits,
  :type,
  keyword_init: true
)

# Request payload for PlayerStat#list.
#
# @!attribute [rw] person_id
#   @return [Integer]
#
# @!attribute [rw] season
#   @return [String, nil]
#
# @!attribute [rw] stat
#   @return [String]
PlayerStatListMatch = Struct.new(
  :person_id,
  :season,
  :stat,
  keyword_init: true
)

# Roster entity data model.
#
# @!attribute [rw] jerseyNumber
#   @return [String, nil]
#
# @!attribute [rw] person
#   @return [Hash, nil]
#
# @!attribute [rw] position
#   @return [Hash, nil]
Roster = Struct.new(
  :jerseyNumber,
  :person,
  :position,
  keyword_init: true
)

# Request payload for Roster#list.
#
# @!attribute [rw] team_id
#   @return [Integer]
#
# @!attribute [rw] season
#   @return [String, nil]
RosterListMatch = Struct.new(
  :team_id,
  :season,
  keyword_init: true
)

# Schedule entity data model.
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] games
#   @return [Array, nil]
#
# @!attribute [rw] totalEvents
#   @return [Integer, nil]
#
# @!attribute [rw] totalGames
#   @return [Integer, nil]
#
# @!attribute [rw] totalItems
#   @return [Integer, nil]
#
# @!attribute [rw] totalMatches
#   @return [Integer, nil]
Schedule = Struct.new(
  :date,
  :games,
  :totalEvents,
  :totalGames,
  :totalItems,
  :totalMatches,
  keyword_init: true
)

# Request payload for Schedule#list.
#
# @!attribute [rw] end_date
#   @return [String, nil]
#
# @!attribute [rw] season
#   @return [String, nil]
#
# @!attribute [rw] start_date
#   @return [String, nil]
#
# @!attribute [rw] team_id
#   @return [Integer, nil]
ScheduleListMatch = Struct.new(
  :end_date,
  :season,
  :start_date,
  :team_id,
  keyword_init: true
)

# Standing entity data model.
#
# @!attribute [rw] conference
#   @return [Hash, nil]
#
# @!attribute [rw] division
#   @return [Hash, nil]
#
# @!attribute [rw] teamRecords
#   @return [Array, nil]
Standing = Struct.new(
  :conference,
  :division,
  :teamRecords,
  keyword_init: true
)

# Request payload for Standing#list.
#
# @!attribute [rw] season
#   @return [String, nil]
StandingListMatch = Struct.new(
  :season,
  keyword_init: true
)

# Team entity data model.
#
# @!attribute [rw] abbreviation
#   @return [String, nil]
#
# @!attribute [rw] conference
#   @return [Hash, nil]
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] division
#   @return [Hash, nil]
#
# @!attribute [rw] firstYearOfPlay
#   @return [String, nil]
#
# @!attribute [rw] franchise
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] locationName
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] teamName
#   @return [String, nil]
#
# @!attribute [rw] teams
#   @return [Array, nil]
#
# @!attribute [rw] venue
#   @return [Hash, nil]
Team = Struct.new(
  :abbreviation,
  :conference,
  :copyright,
  :division,
  :firstYearOfPlay,
  :franchise,
  :id,
  :link,
  :locationName,
  :name,
  :teamName,
  :teams,
  :venue,
  keyword_init: true
)

# Request payload for Team#load.
#
# @!attribute [rw] id
#   @return [Integer]
#
# @!attribute [rw] expand
#   @return [String, nil]
TeamLoadMatch = Struct.new(
  :id,
  :expand,
  keyword_init: true
)

# Request payload for Team#list.
#
# @!attribute [rw] expand
#   @return [String, nil]
#
# @!attribute [rw] season
#   @return [String, nil]
TeamListMatch = Struct.new(
  :expand,
  :season,
  keyword_init: true
)

