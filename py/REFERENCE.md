# NhlApiDocumentation Python SDK Reference

Complete API reference for the NhlApiDocumentation Python SDK.


## NhlApiDocumentationSDK

### Constructor

```python
from nhlapidocumentation_sdk import NhlApiDocumentationSDK

client = NhlApiDocumentationSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NhlApiDocumentationSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = NhlApiDocumentationSDK.test()
```


### Instance Methods

#### `Conference(data=None)`

Create a new `ConferenceEntity` instance. Pass `None` for no initial data.

#### `Division(data=None)`

Create a new `DivisionEntity` instance. Pass `None` for no initial data.

#### `Game(data=None)`

Create a new `GameEntity` instance. Pass `None` for no initial data.

#### `Player(data=None)`

Create a new `PlayerEntity` instance. Pass `None` for no initial data.

#### `PlayerStat(data=None)`

Create a new `PlayerStatEntity` instance. Pass `None` for no initial data.

#### `Roster(data=None)`

Create a new `RosterEntity` instance. Pass `None` for no initial data.

#### `Schedule(data=None)`

Create a new `ScheduleEntity` instance. Pass `None` for no initial data.

#### `Standing(data=None)`

Create a new `StandingEntity` instance. Pass `None` for no initial data.

#### `Team(data=None)`

Create a new `TeamEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## ConferenceEntity

```python
conference = client.Conference()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `list` | No |  |
| `copyright` | `str` | No |  |
| `id` | `int` | No |  |
| `link` | `str` | No |  |
| `name` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Conference().list()
for conference in results:
    print(conference)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Conference().load({"id": "conference_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConferenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DivisionEntity

```python
division = client.Division()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `str` | No |  |
| `division` | `list` | No |  |
| `id` | `int` | No |  |
| `link` | `str` | No |  |
| `name` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Division().list()
for division in results:
    print(division)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Division().load({"id": "division_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DivisionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GameEntity

```python
game = client.Game()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `str` | No |  |
| `game_data` | `dict` | No |  |
| `game_pk` | `int` | No |  |
| `link` | `str` | No |  |
| `live_data` | `dict` | No |  |
| `team` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Game().load({"id": "game_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GameEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlayerEntity

```python
player = client.Player()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `str` | No |  |
| `person` | `list` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Player().load({"id": "player_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlayerEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlayerStatEntity

```python
player_stat = client.PlayerStat()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | `list` | No |  |
| `type` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PlayerStat().list()
for player_stat in results:
    print(player_stat)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlayerStatEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RosterEntity

```python
roster = client.Roster()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | `str` | No |  |
| `person` | `dict` | No |  |
| `position` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Roster().list()
for roster in results:
    print(roster)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RosterEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ScheduleEntity

```python
schedule = client.Schedule()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | `str` | No |  |
| `game` | `list` | No |  |
| `total_event` | `int` | No |  |
| `total_game` | `int` | No |  |
| `total_item` | `int` | No |  |
| `total_match` | `int` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Schedule().list()
for schedule in results:
    print(schedule)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ScheduleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## StandingEntity

```python
standing = client.Standing()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `dict` | No |  |
| `division` | `dict` | No |  |
| `team_record` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Standing().list()
for standing in results:
    print(standing)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StandingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TeamEntity

```python
team = client.Team()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `str` | No |  |
| `conference` | `dict` | No |  |
| `copyright` | `str` | No |  |
| `division` | `dict` | No |  |
| `first_year_of_play` | `str` | No |  |
| `franchise` | `dict` | No |  |
| `id` | `int` | No |  |
| `link` | `str` | No |  |
| `location_name` | `str` | No |  |
| `name` | `str` | No |  |
| `team` | `list` | No |  |
| `team_name` | `str` | No |  |
| `venue` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Team().list()
for team in results:
    print(team)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Team().load({"id": "team_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TeamEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = NhlApiDocumentationSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

