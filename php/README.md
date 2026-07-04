# NhlApiDocumentation PHP SDK



The PHP SDK for the NhlApiDocumentation API — an entity-oriented client using PHP conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases](https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'nhlapidocumentation_sdk.php';

$client = new NhlApiDocumentationSDK();
```

### 2. List conference records

```php
try {
    // list() returns an array of Conference records — iterate directly.
    $conferences = $client->Conference()->list();
    foreach ($conferences as $item) {
        echo $item["id"] . " " . $item["name"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a conference

```php
try {
    // load() returns the bare Conference record (throws on error).
    $conference = $client->Conference()->load(["id" => "example_id"]);
    print_r($conference);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    echo "Error: " . $result["err"]->getMessage();
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = NhlApiDocumentationSDK::test([
    "entity" => ["conference" => ["test01" => ["id" => "test01"]]],
]);

// load() returns the bare mock record (throws on error).
$conference = $client->Conference()->load(["id" => "test01"]);
print_r($conference);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new NhlApiDocumentationSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
NHL_API_DOCUMENTATION_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### NhlApiDocumentationSDK

```php
require_once 'nhlapidocumentation_sdk.php';
$client = new NhlApiDocumentationSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = NhlApiDocumentationSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### NhlApiDocumentationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Conference` | `($data): ConferenceEntity` | Create a Conference entity instance. |
| `Division` | `($data): DivisionEntity` | Create a Division entity instance. |
| `Game` | `($data): GameEntity` | Create a Game entity instance. |
| `Player` | `($data): PlayerEntity` | Create a Player entity instance. |
| `PlayerStat` | `($data): PlayerStatEntity` | Create a PlayerStat entity instance. |
| `Roster` | `($data): RosterEntity` | Create a Roster entity instance. |
| `Schedule` | `($data): ScheduleEntity` | Create a Schedule entity instance. |
| `Standing` | `($data): StandingEntity` | Create a Standing entity instance. |
| `Team` | `($data): TeamEntity` | Create a Team entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `($reqmatch, $ctrl): array` | List entities matching the criteria. |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$conference = $client->Conference();`

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

```php
// load() returns the bare Conference record (throws on error).
$conference = $client->Conference()->load(["id" => "conference_id"]);
```

#### Example: List

```php
// list() returns an array of Conference records (throws on error).
$conferences = $client->Conference()->list();
```


### Division

Create an instance: `$division = $client->Division();`

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

```php
// load() returns the bare Division record (throws on error).
$division = $client->Division()->load(["id" => "division_id"]);
```

#### Example: List

```php
// list() returns an array of Division records (throws on error).
$divisions = $client->Division()->list();
```


### Game

Create an instance: `$game = $client->Game();`

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

```php
// load() returns the bare Game record (throws on error).
$game = $client->Game()->load(["id" => "game_id"]);
```


### Player

Create an instance: `$player = $client->Player();`

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

```php
// load() returns the bare Player record (throws on error).
$player = $client->Player()->load(["id" => "player_id"]);
```


### PlayerStat

Create an instance: `$player_stat = $client->PlayerStat();`

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

```php
// list() returns an array of PlayerStat records (throws on error).
$player_stats = $client->PlayerStat()->list();
```


### Roster

Create an instance: `$roster = $client->Roster();`

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

```php
// list() returns an array of Roster records (throws on error).
$rosters = $client->Roster()->list();
```


### Schedule

Create an instance: `$schedule = $client->Schedule();`

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

```php
// list() returns an array of Schedule records (throws on error).
$schedules = $client->Schedule()->list();
```


### Standing

Create an instance: `$standing = $client->Standing();`

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

```php
// list() returns an array of Standing records (throws on error).
$standings = $client->Standing()->list();
```


### Team

Create an instance: `$team = $client->Team();`

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

```php
// load() returns the bare Team record (throws on error).
$team = $client->Team()->load(["id" => "team_id"]);
```

#### Example: List

```php
// list() returns an array of Team records (throws on error).
$teams = $client->Team()->list();
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
error is returned to the caller as the second element in the return array.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── nhlapidocumentation_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`nhlapidocumentation_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$conference = $client->Conference();
$conference->load(["id" => "example_id"]);

// $conference->dataGet() now returns the loaded conference data
// $conference->matchGet() returns the last match criteria
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
