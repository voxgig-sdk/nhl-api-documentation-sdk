<?php
declare(strict_types=1);

// Typed models for the NhlApiDocumentation SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Conference entity data model. */
class Conference
{
    public ?array $conferences = null;
    public ?string $copyright = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $name = null;
}

/** Request payload for Conference#load. */
class ConferenceLoadMatch
{
    public int $id;
}

/** Request payload for Conference#list. */
class ConferenceListMatch
{
    public ?array $conferences = null;
    public ?string $copyright = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $name = null;
}

/** Division entity data model. */
class Division
{
    public ?string $copyright = null;
    public ?array $divisions = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $name = null;
}

/** Request payload for Division#load. */
class DivisionLoadMatch
{
    public int $id;
}

/** Request payload for Division#list. */
class DivisionListMatch
{
    public ?string $copyright = null;
    public ?array $divisions = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $name = null;
}

/** Game entity data model. */
class Game
{
    public ?array $away = null;
    public ?string $copyright = null;
    public ?array $gameData = null;
    public ?int $gamePk = null;
    public ?array $home = null;
    public ?string $link = null;
    public ?array $liveData = null;
}

/** Request payload for Game#load. */
class GameLoadMatch
{
    public int $id;
}

/** Player entity data model. */
class Player
{
    public ?string $copyright = null;
    public ?array $people = null;
}

/** Request payload for Player#load. */
class PlayerLoadMatch
{
    public int $id;
}

/** PlayerStat entity data model. */
class PlayerStat
{
    public ?array $splits = null;
    public ?array $type = null;
}

/** Request payload for PlayerStat#list. */
class PlayerStatListMatch
{
    public int $person_id;
}

/** Roster entity data model. */
class Roster
{
    public ?string $jerseyNumber = null;
    public ?array $person = null;
    public ?array $position = null;
}

/** Request payload for Roster#list. */
class RosterListMatch
{
    public int $team_id;
}

/** Schedule entity data model. */
class Schedule
{
    public ?string $date = null;
    public ?array $games = null;
    public ?int $totalEvents = null;
    public ?int $totalGames = null;
    public ?int $totalItems = null;
    public ?int $totalMatches = null;
}

/** Request payload for Schedule#list. */
class ScheduleListMatch
{
    public ?string $date = null;
    public ?array $games = null;
    public ?int $totalEvents = null;
    public ?int $totalGames = null;
    public ?int $totalItems = null;
    public ?int $totalMatches = null;
}

/** Standing entity data model. */
class Standing
{
    public ?array $conference = null;
    public ?array $division = null;
    public ?array $teamRecords = null;
}

/** Request payload for Standing#list. */
class StandingListMatch
{
    public ?array $conference = null;
    public ?array $division = null;
    public ?array $teamRecords = null;
}

/** Team entity data model. */
class Team
{
    public ?string $abbreviation = null;
    public ?array $conference = null;
    public ?string $copyright = null;
    public ?array $division = null;
    public ?string $firstYearOfPlay = null;
    public ?array $franchise = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $locationName = null;
    public ?string $name = null;
    public ?string $teamName = null;
    public ?array $teams = null;
    public ?array $venue = null;
}

/** Request payload for Team#load. */
class TeamLoadMatch
{
    public int $id;
}

/** Request payload for Team#list. */
class TeamListMatch
{
    public ?string $abbreviation = null;
    public ?array $conference = null;
    public ?string $copyright = null;
    public ?array $division = null;
    public ?string $firstYearOfPlay = null;
    public ?array $franchise = null;
    public ?int $id = null;
    public ?string $link = null;
    public ?string $locationName = null;
    public ?string $name = null;
    public ?string $teamName = null;
    public ?array $teams = null;
    public ?array $venue = null;
}

