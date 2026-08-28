// Typed models for the NhlApiDocumentation SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/nhl-api-documentation-sdk/go/core"
)

// Conference is the typed data model for the conference entity.
type Conference struct {
	Conferences *[]any `json:"conferences,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ConferenceLoadMatch is the typed request payload for Conference.LoadTyped.
type ConferenceLoadMatch struct {
	Id int `json:"id"`
}

// ConferenceListMatch is the typed request payload for Conference.ListTyped.
type ConferenceListMatch struct {
	Conferences *[]any `json:"conferences,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Division is the typed data model for the division entity.
type Division struct {
	Copyright *string `json:"copyright,omitempty"`
	Divisions *[]any `json:"divisions,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// DivisionLoadMatch is the typed request payload for Division.LoadTyped.
type DivisionLoadMatch struct {
	Id int `json:"id"`
}

// DivisionListMatch is the typed request payload for Division.ListTyped.
type DivisionListMatch struct {
	Copyright *string `json:"copyright,omitempty"`
	Divisions *[]any `json:"divisions,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Game is the typed data model for the game entity.
type Game struct {
	Away *map[string]any `json:"away,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	GameData *map[string]any `json:"gameData,omitempty"`
	GamePk *int `json:"gamePk,omitempty"`
	Home *map[string]any `json:"home,omitempty"`
	Id *string `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	LiveData *map[string]any `json:"liveData,omitempty"`
}

// GameLoadMatch is the typed request payload for Game.LoadTyped.
type GameLoadMatch struct {
	Id int `json:"id"`
}

// Player is the typed data model for the player entity.
type Player struct {
	Copyright *string `json:"copyright,omitempty"`
	Id *string `json:"id,omitempty"`
	People *[]any `json:"people,omitempty"`
}

// PlayerLoadMatch is the typed request payload for Player.LoadTyped.
type PlayerLoadMatch struct {
	Id int `json:"id"`
}

// PlayerStat is the typed data model for the player_stat entity.
type PlayerStat struct {
	Splits *[]any `json:"splits,omitempty"`
	Type *map[string]any `json:"type,omitempty"`
}

// PlayerStatListMatch is the typed request payload for PlayerStat.ListTyped.
type PlayerStatListMatch struct {
	PersonId int `json:"person_id"`
	Season *string `json:"season,omitempty"`
	Stat string `json:"stat"`
}

// Roster is the typed data model for the roster entity.
type Roster struct {
	JerseyNumber *string `json:"jerseyNumber,omitempty"`
	Person *map[string]any `json:"person,omitempty"`
	Position *map[string]any `json:"position,omitempty"`
}

// RosterListMatch is the typed request payload for Roster.ListTyped.
type RosterListMatch struct {
	TeamId int `json:"team_id"`
	Season *string `json:"season,omitempty"`
}

// Schedule is the typed data model for the schedule entity.
type Schedule struct {
	Date *string `json:"date,omitempty"`
	Games *[]any `json:"games,omitempty"`
	TotalEvents *int `json:"totalEvents,omitempty"`
	TotalGames *int `json:"totalGames,omitempty"`
	TotalItems *int `json:"totalItems,omitempty"`
	TotalMatches *int `json:"totalMatches,omitempty"`
}

// ScheduleListMatch is the typed request payload for Schedule.ListTyped.
type ScheduleListMatch struct {
	EndDate *string `json:"end_date,omitempty"`
	Season *string `json:"season,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	TeamId *int `json:"team_id,omitempty"`
}

// Standing is the typed data model for the standing entity.
type Standing struct {
	Conference *map[string]any `json:"conference,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	TeamRecords *[]any `json:"teamRecords,omitempty"`
}

// StandingListMatch is the typed request payload for Standing.ListTyped.
type StandingListMatch struct {
	Season *string `json:"season,omitempty"`
}

// Team is the typed data model for the team entity.
type Team struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	Conference *map[string]any `json:"conference,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	Division *map[string]any `json:"division,omitempty"`
	FirstYearOfPlay *string `json:"firstYearOfPlay,omitempty"`
	Franchise *map[string]any `json:"franchise,omitempty"`
	Id *int `json:"id,omitempty"`
	Link *string `json:"link,omitempty"`
	LocationName *string `json:"locationName,omitempty"`
	Name *string `json:"name,omitempty"`
	TeamName *string `json:"teamName,omitempty"`
	Teams *[]any `json:"teams,omitempty"`
	Venue *map[string]any `json:"venue,omitempty"`
}

// TeamLoadMatch is the typed request payload for Team.LoadTyped.
type TeamLoadMatch struct {
	Id int `json:"id"`
	Expand *string `json:"expand,omitempty"`
}

// TeamListMatch is the typed request payload for Team.ListTyped.
type TeamListMatch struct {
	Expand *string `json:"expand,omitempty"`
	Season *string `json:"season,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
