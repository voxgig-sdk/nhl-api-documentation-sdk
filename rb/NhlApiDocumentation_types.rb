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
# @!attribute [rw] conference
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
  :conference,
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
# @!attribute [rw] conference
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
  :conference,
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
# @!attribute [rw] division
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
  :division,
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
# @!attribute [rw] division
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
  :division,
  :id,
  :link,
  :name,
  keyword_init: true
)

# Game entity data model.
#
# @!attribute [rw] copyright
#   @return [String, nil]
#
# @!attribute [rw] game_data
#   @return [Hash, nil]
#
# @!attribute [rw] game_pk
#   @return [Integer, nil]
#
# @!attribute [rw] link
#   @return [String, nil]
#
# @!attribute [rw] live_data
#   @return [Hash, nil]
#
# @!attribute [rw] team
#   @return [Hash, nil]
Game = Struct.new(
  :copyright,
  :game_data,
  :game_pk,
  :link,
  :live_data,
  :team,
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
# @!attribute [rw] person
#   @return [Array, nil]
Player = Struct.new(
  :copyright,
  :person,
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
# @!attribute [rw] split
#   @return [Array, nil]
#
# @!attribute [rw] type
#   @return [Hash, nil]
PlayerStat = Struct.new(
  :split,
  :type,
  keyword_init: true
)

# Request payload for PlayerStat#list.
#
# @!attribute [rw] person_id
#   @return [Integer]
PlayerStatListMatch = Struct.new(
  :person_id,
  keyword_init: true
)

# Roster entity data model.
#
# @!attribute [rw] jersey_number
#   @return [String, nil]
#
# @!attribute [rw] person
#   @return [Hash, nil]
#
# @!attribute [rw] position
#   @return [Hash, nil]
Roster = Struct.new(
  :jersey_number,
  :person,
  :position,
  keyword_init: true
)

# Request payload for Roster#list.
#
# @!attribute [rw] team_id
#   @return [Integer]
RosterListMatch = Struct.new(
  :team_id,
  keyword_init: true
)

# Schedule entity data model.
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] game
#   @return [Array, nil]
#
# @!attribute [rw] total_event
#   @return [Integer, nil]
#
# @!attribute [rw] total_game
#   @return [Integer, nil]
#
# @!attribute [rw] total_item
#   @return [Integer, nil]
#
# @!attribute [rw] total_match
#   @return [Integer, nil]
Schedule = Struct.new(
  :date,
  :game,
  :total_event,
  :total_game,
  :total_item,
  :total_match,
  keyword_init: true
)

# Request payload for Schedule#list.
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] game
#   @return [Array, nil]
#
# @!attribute [rw] total_event
#   @return [Integer, nil]
#
# @!attribute [rw] total_game
#   @return [Integer, nil]
#
# @!attribute [rw] total_item
#   @return [Integer, nil]
#
# @!attribute [rw] total_match
#   @return [Integer, nil]
ScheduleListMatch = Struct.new(
  :date,
  :game,
  :total_event,
  :total_game,
  :total_item,
  :total_match,
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
# @!attribute [rw] team_record
#   @return [Array, nil]
Standing = Struct.new(
  :conference,
  :division,
  :team_record,
  keyword_init: true
)

# Request payload for Standing#list.
#
# @!attribute [rw] conference
#   @return [Hash, nil]
#
# @!attribute [rw] division
#   @return [Hash, nil]
#
# @!attribute [rw] team_record
#   @return [Array, nil]
StandingListMatch = Struct.new(
  :conference,
  :division,
  :team_record,
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
# @!attribute [rw] first_year_of_play
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
# @!attribute [rw] location_name
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] team
#   @return [Array, nil]
#
# @!attribute [rw] team_name
#   @return [String, nil]
#
# @!attribute [rw] venue
#   @return [Hash, nil]
Team = Struct.new(
  :abbreviation,
  :conference,
  :copyright,
  :division,
  :first_year_of_play,
  :franchise,
  :id,
  :link,
  :location_name,
  :name,
  :team,
  :team_name,
  :venue,
  keyword_init: true
)

# Request payload for Team#load.
#
# @!attribute [rw] id
#   @return [Integer]
TeamLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Team#list.
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
# @!attribute [rw] first_year_of_play
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
# @!attribute [rw] location_name
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] team
#   @return [Array, nil]
#
# @!attribute [rw] team_name
#   @return [String, nil]
#
# @!attribute [rw] venue
#   @return [Hash, nil]
TeamListMatch = Struct.new(
  :abbreviation,
  :conference,
  :copyright,
  :division,
  :first_year_of_play,
  :franchise,
  :id,
  :link,
  :location_name,
  :name,
  :team,
  :team_name,
  :venue,
  keyword_init: true
)

