# NhlApiDocumentation Lua SDK



The Lua SDK for the NhlApiDocumentation API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Conference()` — each with the same small set of operations (`list`, `load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("nhl-api-documentation_sdk")

local client = sdk.new()
```

### 2. List conference records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local conferences, err = client:Conference():list()
if err then error(err) end

for _, item in ipairs(conferences) do
  print(item["id"], item["copyright"])
end
```

### 3. Load a conference

```lua
local conference, err = client:Conference():load({ id = 1 })
if err then error(err) end
print(conference)
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local conferences, err = client:Conference():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Conference():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### NhlApiDocumentationSDK

```lua
local sdk = require("nhl-api-documentation_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NhlApiDocumentationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local conference, err = client:Conference():load({ id = "example_id" })
    if err then error(err) end
    -- conference is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local conference = client:Conference(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | `table` |  |
| `copyright` | `string` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `name` | `string` |  |

#### Example: Load

```lua
local conference, err = client:Conference():load({ id = 1 })
```

#### Example: List

```lua
local conferences, err = client:Conference():list()
```


### Division

Create an instance: `local division = client:Division(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `string` |  |
| `division` | `table` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `name` | `string` |  |

#### Example: Load

```lua
local division, err = client:Division():load({ id = 1 })
```

#### Example: List

```lua
local divisions, err = client:Division():list()
```


### Game

Create an instance: `local game = client:Game(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `string` |  |
| `game_data` | `table` |  |
| `game_pk` | `number` |  |
| `link` | `string` |  |
| `live_data` | `table` |  |
| `team` | `table` |  |

#### Example: Load

```lua
local game, err = client:Game():load({ id = 1 })
```


### Player

Create an instance: `local player = client:Player(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `string` |  |
| `person` | `table` |  |

#### Example: Load

```lua
local player, err = client:Player():load({ id = 1 })
```


### PlayerStat

Create an instance: `local player_stat = client:PlayerStat(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `split` | `table` |  |
| `type` | `table` |  |

#### Example: List

```lua
local player_stats, err = client:PlayerStat():list()
```


### Roster

Create an instance: `local roster = client:Roster(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jersey_number` | `string` |  |
| `person` | `table` |  |
| `position` | `table` |  |

#### Example: List

```lua
local rosters, err = client:Roster():list()
```


### Schedule

Create an instance: `local schedule = client:Schedule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `game` | `table` |  |
| `total_event` | `number` |  |
| `total_game` | `number` |  |
| `total_item` | `number` |  |
| `total_match` | `number` |  |

#### Example: List

```lua
local schedules, err = client:Schedule():list()
```


### Standing

Create an instance: `local standing = client:Standing(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | `table` |  |
| `division` | `table` |  |
| `team_record` | `table` |  |

#### Example: List

```lua
local standings, err = client:Standing():list()
```


### Team

Create an instance: `local team = client:Team(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `abbreviation` | `string` |  |
| `conference` | `table` |  |
| `copyright` | `string` |  |
| `division` | `table` |  |
| `first_year_of_play` | `string` |  |
| `franchise` | `table` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `location_name` | `string` |  |
| `name` | `string` |  |
| `team` | `table` |  |
| `team_name` | `string` |  |
| `venue` | `table` |  |

#### Example: Load

```lua
local team, err = client:Team():load({ id = 1 })
```

#### Example: List

```lua
local teams, err = client:Team():list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── nhl-api-documentation_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`nhl-api-documentation_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local conference = client:Conference()
conference:list()

-- conference:data_get() now returns the conference data from the last list
-- conference:match_get() returns the last match criteria
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
