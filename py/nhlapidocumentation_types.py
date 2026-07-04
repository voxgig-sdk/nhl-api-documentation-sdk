# Typed models for the NhlApiDocumentation SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Conference:
    conference: Optional[list] = None
    copyright: Optional[str] = None
    id: Optional[int] = None
    link: Optional[str] = None
    name: Optional[str] = None


@dataclass
class ConferenceLoadMatch:
    id: int


@dataclass
class ConferenceListMatch:
    conference: Optional[list] = None
    copyright: Optional[str] = None
    id: Optional[int] = None
    link: Optional[str] = None
    name: Optional[str] = None


@dataclass
class Division:
    copyright: Optional[str] = None
    division: Optional[list] = None
    id: Optional[int] = None
    link: Optional[str] = None
    name: Optional[str] = None


@dataclass
class DivisionLoadMatch:
    id: int


@dataclass
class DivisionListMatch:
    copyright: Optional[str] = None
    division: Optional[list] = None
    id: Optional[int] = None
    link: Optional[str] = None
    name: Optional[str] = None


@dataclass
class Game:
    copyright: Optional[str] = None
    game_data: Optional[dict] = None
    game_pk: Optional[int] = None
    link: Optional[str] = None
    live_data: Optional[dict] = None
    team: Optional[dict] = None


@dataclass
class GameLoadMatch:
    id: int


@dataclass
class Player:
    copyright: Optional[str] = None
    person: Optional[list] = None


@dataclass
class PlayerLoadMatch:
    id: int


@dataclass
class PlayerStat:
    split: Optional[list] = None
    type: Optional[dict] = None


@dataclass
class PlayerStatListMatch:
    person_id: int


@dataclass
class Roster:
    jersey_number: Optional[str] = None
    person: Optional[dict] = None
    position: Optional[dict] = None


@dataclass
class RosterListMatch:
    team_id: int


@dataclass
class Schedule:
    date: Optional[str] = None
    game: Optional[list] = None
    total_event: Optional[int] = None
    total_game: Optional[int] = None
    total_item: Optional[int] = None
    total_match: Optional[int] = None


@dataclass
class ScheduleListMatch:
    date: Optional[str] = None
    game: Optional[list] = None
    total_event: Optional[int] = None
    total_game: Optional[int] = None
    total_item: Optional[int] = None
    total_match: Optional[int] = None


@dataclass
class Standing:
    conference: Optional[dict] = None
    division: Optional[dict] = None
    team_record: Optional[list] = None


@dataclass
class StandingListMatch:
    conference: Optional[dict] = None
    division: Optional[dict] = None
    team_record: Optional[list] = None


@dataclass
class Team:
    abbreviation: Optional[str] = None
    conference: Optional[dict] = None
    copyright: Optional[str] = None
    division: Optional[dict] = None
    first_year_of_play: Optional[str] = None
    franchise: Optional[dict] = None
    id: Optional[int] = None
    link: Optional[str] = None
    location_name: Optional[str] = None
    name: Optional[str] = None
    team: Optional[list] = None
    team_name: Optional[str] = None
    venue: Optional[dict] = None


@dataclass
class TeamLoadMatch:
    id: int


@dataclass
class TeamListMatch:
    abbreviation: Optional[str] = None
    conference: Optional[dict] = None
    copyright: Optional[str] = None
    division: Optional[dict] = None
    first_year_of_play: Optional[str] = None
    franchise: Optional[dict] = None
    id: Optional[int] = None
    link: Optional[str] = None
    location_name: Optional[str] = None
    name: Optional[str] = None
    team: Optional[list] = None
    team_name: Optional[str] = None
    venue: Optional[dict] = None

