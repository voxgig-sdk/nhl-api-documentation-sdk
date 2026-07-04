// Typed models for the NhlApiDocumentation SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Conference is the typed data model for the conference entity.
type Conference struct {
	Conference *[]any `json:"conference,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ConferenceLoadMatch is the typed request payload for Conference.LoadTyped.
type ConferenceLoadMatch struct {
	Id int `json:"id"`
}

// ConferenceListMatch mirrors the conference fields as an all-optional match
// filter (Go analog of Partial<Conference>).
type ConferenceListMatch struct {
	Conference *[]any `json:"conference,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Division is the typed data model for the division entity.
type Division struct {
	Copyright *string `json:"copyright,omitempty"`
	Division *[]any `json:"division,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// DivisionLoadMatch is the typed request payload for Division.LoadTyped.
type DivisionLoadMatch struct {
	Id int `json:"id"`
}

// DivisionListMatch mirrors the division fields as an all-optional match
// filter (Go analog of Partial<Division>).
type DivisionListMatch struct {
	Copyright *string `json:"copyright,omitempty"`
	Division *[]any `json:"division,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Game is the typed data model for the game entity.
type Game struct {
	Copyright *string `json:"copyright,omitempty"`
	GameData *map[string]any `json:"game_data,omitempty"`
	GamePk *int `json:"game_pk,omitempty"`
	Link *string `json:"link,omitempty"`
	LiveData *map[string]any `json:"live_data,omitempty"`
	Team *map[string]any `json:"team,omitempty"`
}

// GameLoadMatch is the typed request payload for Game.LoadTyped.
type GameLoadMatch struct {
	Id int `json:"id"`
}

// Player is the typed data model for the player entity.
type Player struct {
	Copyright *string `json:"copyright,omitempty"`
	Person *[]any `json:"person,omitempty"`
}

// PlayerLoadMatch is the typed request payload for Player.LoadTyped.
type PlayerLoadMatch struct {
	Id int `json:"id"`
}

// PlayerStat is the typed data model for the player_stat entity.
type PlayerStat struct {
	Split *[]any `json:"split,omitempty"`
	Type *map[string]any `json:"type,omitempty"`
}

// PlayerStatListMatch is the typed request payload for PlayerStat.ListTyped.
type PlayerStatListMatch struct {
	PersonId int `json:"person_id"`
}

// Roster is the typed data model for the roster entity.
type Roster struct {
	JerseyNumber *string `json:"jersey_number,omitempty"`
	Person *map[string]any `json:"person,omitempty"`
	Position *map[string]any `json:"position,omitempty"`
}

// RosterListMatch is the typed request payload for Roster.ListTyped.
type RosterListMatch struct {
	TeamId int `json:"team_id"`
}

// Schedule is the typed data model for the schedule entity.
type Schedule struct {
	Date *string `json:"date,omitempty"`
	Game *[]any `json:"game,omitempty"`
	TotalEvent *int `json:"total_event,omitempty"`
	TotalGame *int `json:"total_game,omitempty"`
	TotalItem *int `json:"total_item,omitempty"`
	TotalMatch *int `json:"total_match,omitempty"`
}

// ScheduleListMatch mirrors the schedule fields as an all-optional match
// filter (Go analog of Partial<Schedule>).
type ScheduleListMatch struct {
	Date *string `json:"date,omitempty"`
	Game *[]any `json:"game,omitempty"`
	TotalEvent *int `json:"total_event,omitempty"`
	TotalGame *int `json:"total_game,omitempty"`
	TotalItem *int `json:"total_item,omitempty"`
	TotalMatch *int `json:"total_match,omitempty"`
}

// Standing is the typed data model for the standing entity.
type Standing struct {
	Conference *map[string]any `json:"conference,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	TeamRecord *[]any `json:"team_record,omitempty"`
}

// StandingListMatch mirrors the standing fields as an all-optional match
// filter (Go analog of Partial<Standing>).
type StandingListMatch struct {
	Conference *map[string]any `json:"conference,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	TeamRecord *[]any `json:"team_record,omitempty"`
}

// Team is the typed data model for the team entity.
type Team struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	Conference *map[string]any `json:"conference,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	FirstYearOfPlay *string `json:"first_year_of_play,omitempty"`
	Franchise *map[string]any `json:"franchise,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	LocationName *string `json:"location_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Team *[]any `json:"team,omitempty"`
	TeamName *string `json:"team_name,omitempty"`
	Venue *map[string]any `json:"venue,omitempty"`
}

// TeamLoadMatch is the typed request payload for Team.LoadTyped.
type TeamLoadMatch struct {
	Id int `json:"id"`
}

// TeamListMatch mirrors the team fields as an all-optional match
// filter (Go analog of Partial<Team>).
type TeamListMatch struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	Conference *map[string]any `json:"conference,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	FirstYearOfPlay *string `json:"first_year_of_play,omitempty"`
	Franchise *map[string]any `json:"franchise,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	LocationName *string `json:"location_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Team *[]any `json:"team,omitempty"`
	TeamName *string `json:"team_name,omitempty"`
	Venue *map[string]any `json:"venue,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
