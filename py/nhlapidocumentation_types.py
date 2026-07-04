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
    conference: list
    copyright: str
    id: int
    link: str
    name: str


class ConferenceLoadMatch(TypedDict):
    id: int


class ConferenceListMatch(TypedDict, total=False):
    conference: list
    copyright: str
    id: int
    link: str
    name: str


class Division(TypedDict, total=False):
    copyright: str
    division: list
    id: int
    link: str
    name: str


class DivisionLoadMatch(TypedDict):
    id: int


class DivisionListMatch(TypedDict, total=False):
    copyright: str
    division: list
    id: int
    link: str
    name: str


class Game(TypedDict, total=False):
    copyright: str
    game_data: dict
    game_pk: int
    link: str
    live_data: dict
    team: dict


class GameLoadMatch(TypedDict):
    id: int


class Player(TypedDict, total=False):
    copyright: str
    person: list


class PlayerLoadMatch(TypedDict):
    id: int


class PlayerStat(TypedDict, total=False):
    split: list
    type: dict


class PlayerStatListMatch(TypedDict):
    person_id: int


class Roster(TypedDict, total=False):
    jersey_number: str
    person: dict
    position: dict


class RosterListMatch(TypedDict):
    team_id: int


class Schedule(TypedDict, total=False):
    date: str
    game: list
    total_event: int
    total_game: int
    total_item: int
    total_match: int


class ScheduleListMatch(TypedDict, total=False):
    date: str
    game: list
    total_event: int
    total_game: int
    total_item: int
    total_match: int


class Standing(TypedDict, total=False):
    conference: dict
    division: dict
    team_record: list


class StandingListMatch(TypedDict, total=False):
    conference: dict
    division: dict
    team_record: list


class Team(TypedDict, total=False):
    abbreviation: str
    conference: dict
    copyright: str
    division: dict
    first_year_of_play: str
    franchise: dict
    id: int
    link: str
    location_name: str
    name: str
    team: list
    team_name: str
    venue: dict


class TeamLoadMatch(TypedDict):
    id: int


class TeamListMatch(TypedDict, total=False):
    abbreviation: str
    conference: dict
    copyright: str
    division: dict
    first_year_of_play: str
    franchise: dict
    id: int
    link: str
    location_name: str
    name: str
    team: list
    team_name: str
    venue: dict
