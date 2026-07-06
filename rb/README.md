# NhlApiDocumentation Ruby SDK



The Ruby SDK for the NhlApiDocumentation API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Conference` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases](https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "NhlApiDocumentation_sdk"

client = NhlApiDocumentationSDK.new
```

### 2. List conference records

```ruby
begin
  # list returns an Array of Conference records — iterate directly.
  conferences = client.Conference.list
  conferences.each do |item|
    puts "#{item["id"]} #{item["conference"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a conference

```ruby
begin
  # load returns the bare Conference record (raises on error).
  conference = client.Conference.load({ "id" => "example_id" })
  puts conference
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  conferences = client.Conference.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = NhlApiDocumentationSDK.test({
  "entity" => { "conference" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the bare mock record (raises on error).
conference = client.Conference.list()
puts conference
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = NhlApiDocumentationSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### NhlApiDocumentationSDK

```ruby
require_relative "NhlApiDocumentation_sdk"
client = NhlApiDocumentationSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = NhlApiDocumentationSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NhlApiDocumentationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `NhlApiDocumentationError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Create an instance: `conference = client.Conference`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | `Array` |  |
| `copyright` | `String` |  |
| `id` | `Integer` |  |
| `link` | `String` |  |
| `name` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Conference record (raises on error).
conference = client.Conference.load({ "id" => "conference_id" })
```

#### Example: List

```ruby
# list returns an Array of Conference records (raises on error).
conferences = client.Conference.list
```


### Division

Create an instance: `division = client.Division`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `String` |  |
| `division` | `Array` |  |
| `id` | `Integer` |  |
| `link` | `String` |  |
| `name` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Division record (raises on error).
division = client.Division.load({ "id" => "division_id" })
```

#### Example: List

```ruby
# list returns an Array of Division records (raises on error).
divisions = client.Division.list
```


### Game

Create an instance: `game = client.Game`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `String` |  |
| `game_data` | `Hash` |  |
| `game_pk` | `Integer` |  |
| `link` | `String` |  |
| `live_data` | `Hash` |  |
| `team` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare Game record (raises on error).
game = client.Game.load({ "id" => "game_id" })
```


### Player

Create an instance: `player = client.Player`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `copyright` | `String` |  |
| `person` | `Array` |  |

#### Example: Load

```ruby
# load returns the bare Player record (raises on error).
player = client.Player.load({ "id" => "player_id" })
```


### PlayerStat

Create an instance: `player_stat = client.PlayerStat`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `split` | `Array` |  |
| `type` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of PlayerStat records (raises on error).
player_stats = client.PlayerStat.list
```


### Roster

Create an instance: `roster = client.Roster`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jersey_number` | `String` |  |
| `person` | `Hash` |  |
| `position` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of Roster records (raises on error).
rosters = client.Roster.list
```


### Schedule

Create an instance: `schedule = client.Schedule`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `date` | `String` |  |
| `game` | `Array` |  |
| `total_event` | `Integer` |  |
| `total_game` | `Integer` |  |
| `total_item` | `Integer` |  |
| `total_match` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of Schedule records (raises on error).
schedules = client.Schedule.list
```


### Standing

Create an instance: `standing = client.Standing`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `conference` | `Hash` |  |
| `division` | `Hash` |  |
| `team_record` | `Array` |  |

#### Example: List

```ruby
# list returns an Array of Standing records (raises on error).
standings = client.Standing.list
```


### Team

Create an instance: `team = client.Team`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `abbreviation` | `String` |  |
| `conference` | `Hash` |  |
| `copyright` | `String` |  |
| `division` | `Hash` |  |
| `first_year_of_play` | `String` |  |
| `franchise` | `Hash` |  |
| `id` | `Integer` |  |
| `link` | `String` |  |
| `location_name` | `String` |  |
| `name` | `String` |  |
| `team` | `Array` |  |
| `team_name` | `String` |  |
| `venue` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare Team record (raises on error).
team = client.Team.load({ "id" => "team_id" })
```

#### Example: List

```ruby
# list returns an Array of Team records (raises on error).
teams = client.Team.list
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── NhlApiDocumentation_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`NhlApiDocumentation_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
conference = client.Conference
conference.list()

# conference.data_get now returns the conference data from the last list
# conference.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
