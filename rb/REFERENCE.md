# NhlApiDocumentation Ruby SDK Reference

Complete API reference for the NhlApiDocumentation Ruby SDK.


## NhlApiDocumentationSDK

### Constructor

```ruby
require_relative 'NhlApiDocumentation_sdk'

client = NhlApiDocumentationSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NhlApiDocumentationSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = NhlApiDocumentationSDK.test
```


### Instance Methods

#### `Conference(data = nil)`

Create a new `Conference` entity instance. Pass `nil` for no initial data.

#### `Division(data = nil)`

Create a new `Division` entity instance. Pass `nil` for no initial data.

#### `Game(data = nil)`

Create a new `Game` entity instance. Pass `nil` for no initial data.

#### `Player(data = nil)`

Create a new `Player` entity instance. Pass `nil` for no initial data.

#### `PlayerStat(data = nil)`

Create a new `PlayerStat` entity instance. Pass `nil` for no initial data.

#### `Roster(data = nil)`

Create a new `Roster` entity instance. Pass `nil` for no initial data.

#### `Schedule(data = nil)`

Create a new `Schedule` entity instance. Pass `nil` for no initial data.

#### `Standing(data = nil)`

Create a new `Standing` entity instance. Pass `nil` for no initial data.

#### `Team(data = nil)`

Create a new `Team` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## ConferenceEntity

```ruby
conference = client.Conference
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `Array` | No |  |
| `copyright` | `String` | No |  |
| `id` | `Integer` | No |  |
| `link` | `String` | No |  |
| `name` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Conference.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Conference.load({ "id" => "conference_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ConferenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DivisionEntity

```ruby
division = client.Division
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `String` | No |  |
| `division` | `Array` | No |  |
| `id` | `Integer` | No |  |
| `link` | `String` | No |  |
| `name` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Division.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Division.load({ "id" => "division_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DivisionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GameEntity

```ruby
game = client.Game
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `String` | No |  |
| `game_data` | `Hash` | No |  |
| `game_pk` | `Integer` | No |  |
| `link` | `String` | No |  |
| `live_data` | `Hash` | No |  |
| `team` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Game.load({ "id" => "game_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GameEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlayerEntity

```ruby
player = client.Player
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `String` | No |  |
| `person` | `Array` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Player.load({ "id" => "player_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlayerEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlayerStatEntity

```ruby
player_stat = client.PlayerStat
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | `Array` | No |  |
| `type` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PlayerStat.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlayerStatEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RosterEntity

```ruby
roster = client.Roster
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | `String` | No |  |
| `person` | `Hash` | No |  |
| `position` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Roster.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RosterEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ScheduleEntity

```ruby
schedule = client.Schedule
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | `String` | No |  |
| `game` | `Array` | No |  |
| `total_event` | `Integer` | No |  |
| `total_game` | `Integer` | No |  |
| `total_item` | `Integer` | No |  |
| `total_match` | `Integer` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Schedule.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ScheduleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## StandingEntity

```ruby
standing = client.Standing
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `Hash` | No |  |
| `division` | `Hash` | No |  |
| `team_record` | `Array` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Standing.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `StandingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TeamEntity

```ruby
team = client.Team
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `String` | No |  |
| `conference` | `Hash` | No |  |
| `copyright` | `String` | No |  |
| `division` | `Hash` | No |  |
| `first_year_of_play` | `String` | No |  |
| `franchise` | `Hash` | No |  |
| `id` | `Integer` | No |  |
| `link` | `String` | No |  |
| `location_name` | `String` | No |  |
| `name` | `String` | No |  |
| `team` | `Array` | No |  |
| `team_name` | `String` | No |  |
| `venue` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Team.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Team.load({ "id" => "team_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TeamEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = NhlApiDocumentationSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

