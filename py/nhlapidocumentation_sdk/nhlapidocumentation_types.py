# Typed models for the NhlApiDocumentation SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Conference(TypedDict, total=False):
    conferences: list
    copyright: str
    id: int
    link: str
    name: str


class ConferenceLoadMatch(TypedDict):
    id: int


class ConferenceListMatch(TypedDict, total=False):
    conferences: list
    copyright: str
    id: int
    link: str
    name: str


class Division(TypedDict, total=False):
    copyright: str
    divisions: list
    id: int
    link: str
    name: str


class DivisionLoadMatch(TypedDict):
    id: int


class DivisionListMatch(TypedDict, total=False):
    copyright: str
    divisions: list
    id: int
    link: str
    name: str


class Game(TypedDict, total=False):
    away: dict
    copyright: str
    gameData: dict
    gamePk: int
    home: dict
    id: str
    link: str
    liveData: dict


class GameLoadMatch(TypedDict):
    id: int


class Player(TypedDict, total=False):
    copyright: str
    id: str
    people: list


class PlayerLoadMatch(TypedDict):
    id: int


class PlayerStat(TypedDict, total=False):
    splits: list
    type: dict


class PlayerStatListMatchRequired(TypedDict):
    person_id: int
    stat: str


class PlayerStatListMatch(PlayerStatListMatchRequired, total=False):
    season: str


class Roster(TypedDict, total=False):
    jerseyNumber: str
    person: dict
    position: dict


class RosterListMatchRequired(TypedDict):
    team_id: int


class RosterListMatch(RosterListMatchRequired, total=False):
    season: str


class Schedule(TypedDict, total=False):
    date: str
    games: list
    totalEvents: int
    totalGames: int
    totalItems: int
    totalMatches: int


class ScheduleListMatch(TypedDict, total=False):
    end_date: str
    season: str
    start_date: str
    team_id: int


class Standing(TypedDict, total=False):
    conference: dict
    division: dict
    teamRecords: list


class StandingListMatch(TypedDict, total=False):
    season: str


class Team(TypedDict, total=False):
    abbreviation: str
    conference: dict
    copyright: str
    division: dict
    firstYearOfPlay: str
    franchise: dict
    id: int
    link: str
    locationName: str
    name: str
    teamName: str
    teams: list
    venue: dict


class TeamLoadMatchRequired(TypedDict):
    id: int


class TeamLoadMatch(TeamLoadMatchRequired, total=False):
    expand: str


class TeamListMatch(TypedDict, total=False):
    expand: str
    season: str
