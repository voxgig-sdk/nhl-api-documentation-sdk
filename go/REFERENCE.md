# NhlApiDocumentation Golang SDK Reference

Complete API reference for the NhlApiDocumentation Golang SDK.


## NhlApiDocumentationSDK

### Constructor

```go
func NewNhlApiDocumentationSDK(options map[string]any) *NhlApiDocumentationSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *NhlApiDocumentationSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *NhlApiDocumentationSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Conference(data map[string]any) NhlApiDocumentationEntity`

Create a new `Conference` entity instance. Pass `nil` for no initial data.

#### `Division(data map[string]any) NhlApiDocumentationEntity`

Create a new `Division` entity instance. Pass `nil` for no initial data.

#### `Game(data map[string]any) NhlApiDocumentationEntity`

Create a new `Game` entity instance. Pass `nil` for no initial data.

#### `Player(data map[string]any) NhlApiDocumentationEntity`

Create a new `Player` entity instance. Pass `nil` for no initial data.

#### `PlayerStat(data map[string]any) NhlApiDocumentationEntity`

Create a new `PlayerStat` entity instance. Pass `nil` for no initial data.

#### `Roster(data map[string]any) NhlApiDocumentationEntity`

Create a new `Roster` entity instance. Pass `nil` for no initial data.

#### `Schedule(data map[string]any) NhlApiDocumentationEntity`

Create a new `Schedule` entity instance. Pass `nil` for no initial data.

#### `Standing(data map[string]any) NhlApiDocumentationEntity`

Create a new `Standing` entity instance. Pass `nil` for no initial data.

#### `Team(data map[string]any) NhlApiDocumentationEntity`

Create a new `Team` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## ConferenceEntity

```go
conference := client.Conference(nil)
fmt.Println(conference.GetName()) // "conference"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `[]any` | No |  |
| `copyright` | `string` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Conference(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Conference(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ConferenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DivisionEntity

```go
division := client.Division(nil)
fmt.Println(division.GetName()) // "division"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `division` | `[]any` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Division(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Division(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DivisionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GameEntity

```go
game := client.Game(nil)
fmt.Println(game.GetName()) // "game"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `game_data` | `map[string]any` | No |  |
| `game_pk` | `int` | No |  |
| `link` | `string` | No |  |
| `live_data` | `map[string]any` | No |  |
| `team` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Game(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GameEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlayerEntity

```go
player := client.Player(nil)
fmt.Println(player.GetName()) // "player"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `person` | `[]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Player(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlayerEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlayerStatEntity

```go
playerStat := client.PlayerStat(nil)
fmt.Println(playerStat.GetName()) // "player_stat"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | `[]any` | No |  |
| `type` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PlayerStat(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlayerStatEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RosterEntity

```go
roster := client.Roster(nil)
fmt.Println(roster.GetName()) // "roster"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | `string` | No |  |
| `person` | `map[string]any` | No |  |
| `position` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Roster(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RosterEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ScheduleEntity

```go
schedule := client.Schedule(nil)
fmt.Println(schedule.GetName()) // "schedule"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | `string` | No |  |
| `game` | `[]any` | No |  |
| `total_event` | `int` | No |  |
| `total_game` | `int` | No |  |
| `total_item` | `int` | No |  |
| `total_match` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Schedule(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ScheduleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## StandingEntity

```go
standing := client.Standing(nil)
fmt.Println(standing.GetName()) // "standing"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `map[string]any` | No |  |
| `division` | `map[string]any` | No |  |
| `team_record` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Standing(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `StandingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TeamEntity

```go
team := client.Team(nil)
fmt.Println(team.GetName()) // "team"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No |  |
| `conference` | `map[string]any` | No |  |
| `copyright` | `string` | No |  |
| `division` | `map[string]any` | No |  |
| `first_year_of_play` | `string` | No |  |
| `franchise` | `map[string]any` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `location_name` | `string` | No |  |
| `name` | `string` | No |  |
| `team` | `[]any` | No |  |
| `team_name` | `string` | No |  |
| `venue` | `map[string]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Team(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Team(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TeamEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewNhlApiDocumentationSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

