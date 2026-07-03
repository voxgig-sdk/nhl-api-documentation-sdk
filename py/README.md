# NhlApiDocumentation Python SDK



The Python SDK for the NhlApiDocumentation API — an entity-oriented client following Pythonic conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
pip install voxgig-sdk-nhl-api-documentation
```

Or install from source:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from nhlapidocumentation_sdk import NhlApiDocumentationSDK

client = NhlApiDocumentationSDK({
    "apikey": os.environ.get("NHL-API-DOCUMENTATION_APIKEY"),
})
```

### 2. List conferences

```python
result, err = client.Conference().list()
if err:
    raise Exception(err)

if isinstance(result, list):
    for item in result:
        d = item.data_get()
        print(d["id"], d["name"])
```

### 3. Load a conference

```python
result, err = client.Conference().load({"id": "example_id"})
if err:
    raise Exception(err)
print(result)
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
if err:
    raise Exception(err)

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
```

### Prepare a request without sending it

```python
fetchdef, err = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})
if err:
    raise Exception(err)

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = NhlApiDocumentationSDK.test()

result, err = client.NhlApiDocumentation().load({"id": "test01"})
# result contains mock response data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = NhlApiDocumentationSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NHL-API-DOCUMENTATION_TEST_LIVE=TRUE
NHL-API-DOCUMENTATION_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### NhlApiDocumentationSDK

```python
from nhlapidocumentation_sdk import NhlApiDocumentationSDK

client = NhlApiDocumentationSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = NhlApiDocumentationSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### NhlApiDocumentationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> (dict, err)` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> (dict, err)` | Build and send an HTTP request. |
| `Conference` | `(data) -> ConferenceEntity` | Create a Conference entity instance. |
| `Division` | `(data) -> DivisionEntity` | Create a Division entity instance. |
| `Game` | `(data) -> GameEntity` | Create a Game entity instance. |
| `Player` | `(data) -> PlayerEntity` | Create a Player entity instance. |
| `PlayerStat` | `(data) -> PlayerStatEntity` | Create a PlayerStat entity instance. |
| `Roster` | `(data) -> RosterEntity` | Create a Roster entity instance. |
| `Schedule` | `(data) -> ScheduleEntity` | Create a Schedule entity instance. |
| `Standing` | `(data) -> StandingEntity` | Create a Standing entity instance. |
| `Team` | `(data) -> TeamEntity` | Create a Team entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> (any, err)` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> (any, err)` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> (any, err)` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> (any, err)` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> (any, err)` | Remove an entity. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return `(any, err)`. The first value is a
`dict` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### Conference

| Field | Description |
| --- | --- |
| `conference` |  |
| `copyright` |  |
| `id` |  |
| `link` |  |
| `name` |  |

Operations: List, Load.

API path: `/conferences`

#### Division

| Field | Description |
| --- | --- |
| `copyright` |  |
| `division` |  |
| `id` |  |
| `link` |  |
| `name` |  |

Operations: List, Load.

API path: `/divisions`

#### Game

| Field | Description |
| --- | --- |
| `copyright` |  |
| `game_data` |  |
| `game_pk` |  |
| `link` |  |
| `live_data` |  |
| `team` |  |

Operations: Load.

API path: `/game/{id}/boxscore`

#### Player

| Field | Description |
| --- | --- |
| `copyright` |  |
| `person` |  |

Operations: Load.

API path: `/people/{id}`

#### PlayerStat

| Field | Description |
| --- | --- |
| `split` |  |
| `type` |  |

Operations: List.

API path: `/people/{id}/stats`

#### Roster

| Field | Description |
| --- | --- |
| `jersey_number` |  |
| `person` |  |
| `position` |  |

Operations: List.

API path: `/teams/{id}/roster`

#### Schedule

| Field | Description |
| --- | --- |
| `date` |  |
| `game` |  |
| `total_event` |  |
| `total_game` |  |
| `total_item` |  |
| `total_match` |  |

Operations: List.

API path: `/schedule`

#### Standing

| Field | Description |
| --- | --- |
| `conference` |  |
| `division` |  |
| `team_record` |  |

Operations: List.

API path: `/standings`

#### Team

| Field | Description |
| --- | --- |
| `abbreviation` |  |
| `conference` |  |
| `copyright` |  |
| `division` |  |
| `first_year_of_play` |  |
| `franchise` |  |
| `id` |  |
| `link` |  |
| `location_name` |  |
| `name` |  |
| `team` |  |
| `team_name` |  |
| `venue` |  |

Operations: List, Load.

API path: `/teams`



## Entities


### Conference

Create an instance: `const conference = client.Conference()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | ``$ARRAY`` |  |
| `copyright` | ``$STRING`` |  |
| `id` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |

#### Example: Load

```ts
const conference = await client.Conference().load({ id: 'conference_id' })
```

#### Example: List

```ts
const conferences = await client.Conference().list()
```


### Division

Create an instance: `const division = client.Division()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | ``$STRING`` |  |
| `division` | ``$ARRAY`` |  |
| `id` | ``$INTEGER`` |  |
| `link` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |

#### Example: Load

```ts
const division = await client.Division().load({ id: 'division_id' })
```

#### Example: List

```ts
const divisions = await client.Division().list()
```


### Game

Create an instance: `const game = client.Game()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

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

```ts
const game = await client.Game().load({ id: 'game_id' })
```


### Player

Create an instance: `const player = client.Player()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | ``$STRING`` |  |
| `person` | ``$ARRAY`` |  |

#### Example: Load

```ts
const player = await client.Player().load({ id: 'player_id' })
```


### PlayerStat

Create an instance: `const player_stat = client.PlayerStat()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `split` | ``$ARRAY`` |  |
| `type` | ``$OBJECT`` |  |

#### Example: List

```ts
const player_stats = await client.PlayerStat().list()
```


### Roster

Create an instance: `const roster = client.Roster()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jersey_number` | ``$STRING`` |  |
| `person` | ``$OBJECT`` |  |
| `position` | ``$OBJECT`` |  |

#### Example: List

```ts
const rosters = await client.Roster().list()
```


### Schedule

Create an instance: `const schedule = client.Schedule()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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

```ts
const schedules = await client.Schedule().list()
```


### Standing

Create an instance: `const standing = client.Standing()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | ``$OBJECT`` |  |
| `division` | ``$OBJECT`` |  |
| `team_record` | ``$ARRAY`` |  |

#### Example: List

```ts
const standings = await client.Standing().list()
```


### Team

Create an instance: `const team = client.Team()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

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

```ts
const team = await client.Team().load({ id: 'team_id' })
```

#### Example: List

```ts
const teams = await client.Team().list()
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
error is returned to the caller as the second element in the return tuple.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── nhlapidocumentation_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`nhlapidocumentation_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
moon = client.Moon()
moon.load({"planet_id": "earth", "id": "luna"})

# moon.data_get() now returns the loaded moon data
# moon.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
