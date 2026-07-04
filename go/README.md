# NhlApiDocumentation Golang SDK



The Golang SDK for the NhlApiDocumentation API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/nhl-api-documentation-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/nhl-api-documentation-sdk/go=../nhl-api-documentation-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/nhl-api-documentation-sdk/go"
)

func main() {
    client := sdk.New()

    // List conference records — the value is the array of records itself.
    conferences, err := client.Conference(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range conferences.([]any) {
        fmt.Println(item)
    }

    // Load a single conference — the value is the loaded record.
    conference, err := client.Conference(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(conference)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

conference, err := client.Conference(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(conference) // the loaded mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewNhlApiDocumentationSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NHL_API_DOCUMENTATION_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewNhlApiDocumentationSDK

```go
func NewNhlApiDocumentationSDK(options map[string]any) *NhlApiDocumentationSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *NhlApiDocumentationSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NhlApiDocumentationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Conference` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Conference entity instance. |
| `Division` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Division entity instance. |
| `Game` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Game entity instance. |
| `Player` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Player entity instance. |
| `PlayerStat` | `(data map[string]any) NhlApiDocumentationEntity` | Create a PlayerStat entity instance. |
| `Roster` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Roster entity instance. |
| `Schedule` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Schedule entity instance. |
| `Standing` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Standing entity instance. |
| `Team` | `(data map[string]any) NhlApiDocumentationEntity` | Create a Team entity instance. |

### Entity interface (NhlApiDocumentationEntity)

All entities implement the `NhlApiDocumentationEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    conference, err := client.Conference(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil { /* handle */ }
    // conference is the loaded record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Conference

| Field | Description |
| --- | --- |
| `"conference"` |  |
| `"copyright"` |  |
| `"id"` |  |
| `"link"` |  |
| `"name"` |  |

Operations: List, Load.

API path: `/conferences`

#### Division

| Field | Description |
| --- | --- |
| `"copyright"` |  |
| `"division"` |  |
| `"id"` |  |
| `"link"` |  |
| `"name"` |  |

Operations: List, Load.

API path: `/divisions`

#### Game

| Field | Description |
| --- | --- |
| `"copyright"` |  |
| `"game_data"` |  |
| `"game_pk"` |  |
| `"link"` |  |
| `"live_data"` |  |
| `"team"` |  |

Operations: Load.

API path: `/game/{id}/boxscore`

#### Player

| Field | Description |
| --- | --- |
| `"copyright"` |  |
| `"person"` |  |

Operations: Load.

API path: `/people/{id}`

#### PlayerStat

| Field | Description |
| --- | --- |
| `"split"` |  |
| `"type"` |  |

Operations: List.

API path: `/people/{id}/stats`

#### Roster

| Field | Description |
| --- | --- |
| `"jersey_number"` |  |
| `"person"` |  |
| `"position"` |  |

Operations: List.

API path: `/teams/{id}/roster`

#### Schedule

| Field | Description |
| --- | --- |
| `"date"` |  |
| `"game"` |  |
| `"total_event"` |  |
| `"total_game"` |  |
| `"total_item"` |  |
| `"total_match"` |  |

Operations: List.

API path: `/schedule`

#### Standing

| Field | Description |
| --- | --- |
| `"conference"` |  |
| `"division"` |  |
| `"team_record"` |  |

Operations: List.

API path: `/standings`

#### Team

| Field | Description |
| --- | --- |
| `"abbreviation"` |  |
| `"conference"` |  |
| `"copyright"` |  |
| `"division"` |  |
| `"first_year_of_play"` |  |
| `"franchise"` |  |
| `"id"` |  |
| `"link"` |  |
| `"location_name"` |  |
| `"name"` |  |
| `"team"` |  |
| `"team_name"` |  |
| `"venue"` |  |

Operations: List, Load.

API path: `/teams`



## Entities


### Conference

Create an instance: `conference := client.Conference(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | ``$ARRAY`` |  |
| `copyright` | ``$STRING`` |  |
| `id` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |

#### Example: Load

```go
conference, err := client.Conference(nil).Load(map[string]any{"id": "conference_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(conference) // the loaded record
```

#### Example: List

```go
conferences, err := client.Conference(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(conferences) // the array of records
```


### Division

Create an instance: `division := client.Division(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | ``$STRING`` |  |
| `division` | ``$ARRAY`` |  |
| `id` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |

#### Example: Load

```go
division, err := client.Division(nil).Load(map[string]any{"id": "division_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(division) // the loaded record
```

#### Example: List

```go
divisions, err := client.Division(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(divisions) // the array of records
```


### Game

Create an instance: `game := client.Game(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | ``$STRING`` |  |
| `game_data` | ``$OBJECT`` |  |
| `game_pk` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `live_data` | ``$OBJECT`` |  |
| `team` | ``$OBJECT`` |  |

#### Example: Load

```go
game, err := client.Game(nil).Load(map[string]any{"id": "game_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(game) // the loaded record
```


### Player

Create an instance: `player := client.Player(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | ``$STRING`` |  |
| `person` | ``$ARRAY`` |  |

#### Example: Load

```go
player, err := client.Player(nil).Load(map[string]any{"id": "player_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(player) // the loaded record
```


### PlayerStat

Create an instance: `player_stat := client.PlayerStat(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `split` | ``$ARRAY`` |  |
| `type` | ``$OBJECT`` |  |

#### Example: List

```go
player_stats, err := client.PlayerStat(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(player_stats) // the array of records
```


### Roster

Create an instance: `roster := client.Roster(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jersey_number` | ``$STRING`` |  |
| `person` | ``$OBJECT`` |  |
| `position` | ``$OBJECT`` |  |

#### Example: List

```go
rosters, err := client.Roster(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(rosters) // the array of records
```


### Schedule

Create an instance: `schedule := client.Schedule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date` | ``$STRING`` |  |
| `game` | ``$ARRAY`` |  |
| `total_event` | ``$INTEGER`` |  |
| `total_game` | ``$INTEGER`` |  |
| `total_item` | ``$INTEGER`` |  |
| `total_match` | ``$INTEGER`` |  |

#### Example: List

```go
schedules, err := client.Schedule(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(schedules) // the array of records
```


### Standing

Create an instance: `standing := client.Standing(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | ``$OBJECT`` |  |
| `division` | ``$OBJECT`` |  |
| `team_record` | ``$ARRAY`` |  |

#### Example: List

```go
standings, err := client.Standing(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(standings) // the array of records
```


### Team

Create an instance: `team := client.Team(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `abbreviation` | ``$STRING`` |  |
| `conference` | ``$OBJECT`` |  |
| `copyright` | ``$STRING`` |  |
| `division` | ``$OBJECT`` |  |
| `first_year_of_play` | ``$STRING`` |  |
| `franchise` | ``$OBJECT`` |  |
| `id` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `location_name` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |
| `team` | ``$ARRAY`` |  |
| `team_name` | ``$STRING`` |  |
| `venue` | ``$OBJECT`` |  |

#### Example: Load

```go
team, err := client.Team(nil).Load(map[string]any{"id": "team_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(team) // the loaded record
```

#### Example: List

```go
teams, err := client.Team(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(teams) // the array of records
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/nhl-api-documentation-sdk/go/
├── nhl-api-documentation.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/nhl-api-documentation-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
conference := client.Conference(nil)
conference.Load(map[string]any{"id": "example_id"}, nil)

// conference.Data() now returns the loaded conference data
// conference.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
