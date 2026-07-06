# NhlApiDocumentation PHP SDK Reference

Complete API reference for the NhlApiDocumentation PHP SDK.


## NhlApiDocumentationSDK

### Constructor

```php
require_once __DIR__ . '/nhlapidocumentation_sdk.php';

$client = new NhlApiDocumentationSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NhlApiDocumentationSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = NhlApiDocumentationSDK::test();
```


### Instance Methods

#### `Conference($data = null)`

Create a new `ConferenceEntity` instance. Pass `null` for no initial data.

#### `Division($data = null)`

Create a new `DivisionEntity` instance. Pass `null` for no initial data.

#### `Game($data = null)`

Create a new `GameEntity` instance. Pass `null` for no initial data.

#### `Player($data = null)`

Create a new `PlayerEntity` instance. Pass `null` for no initial data.

#### `PlayerStat($data = null)`

Create a new `PlayerStatEntity` instance. Pass `null` for no initial data.

#### `Roster($data = null)`

Create a new `RosterEntity` instance. Pass `null` for no initial data.

#### `Schedule($data = null)`

Create a new `ScheduleEntity` instance. Pass `null` for no initial data.

#### `Standing($data = null)`

Create a new `StandingEntity` instance. Pass `null` for no initial data.

#### `Team($data = null)`

Create a new `TeamEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): NhlApiDocumentationUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## ConferenceEntity

```php
$conference = $client->Conference();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `array` | No |  |
| `copyright` | `string` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Conference()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Conference()->load(["id" => "conference_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ConferenceEntity`

Create a new `ConferenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DivisionEntity

```php
$division = $client->Division();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `division` | `array` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Division()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Division()->load(["id" => "division_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DivisionEntity`

Create a new `DivisionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GameEntity

```php
$game = $client->Game();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `game_data` | `array` | No |  |
| `game_pk` | `int` | No |  |
| `link` | `string` | No |  |
| `live_data` | `array` | No |  |
| `team` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Game()->load(["id" => "game_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GameEntity`

Create a new `GameEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlayerEntity

```php
$player = $client->Player();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `person` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Player()->load(["id" => "player_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlayerEntity`

Create a new `PlayerEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlayerStatEntity

```php
$player_stat = $client->PlayerStat();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | `array` | No |  |
| `type` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PlayerStat()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlayerStatEntity`

Create a new `PlayerStatEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RosterEntity

```php
$roster = $client->Roster();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | `string` | No |  |
| `person` | `array` | No |  |
| `position` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Roster()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RosterEntity`

Create a new `RosterEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ScheduleEntity

```php
$schedule = $client->Schedule();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | `string` | No |  |
| `game` | `array` | No |  |
| `total_event` | `int` | No |  |
| `total_game` | `int` | No |  |
| `total_item` | `int` | No |  |
| `total_match` | `int` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Schedule()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ScheduleEntity`

Create a new `ScheduleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## StandingEntity

```php
$standing = $client->Standing();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `array` | No |  |
| `division` | `array` | No |  |
| `team_record` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Standing()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): StandingEntity`

Create a new `StandingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TeamEntity

```php
$team = $client->Team();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No |  |
| `conference` | `array` | No |  |
| `copyright` | `string` | No |  |
| `division` | `array` | No |  |
| `first_year_of_play` | `string` | No |  |
| `franchise` | `array` | No |  |
| `id` | `int` | No |  |
| `link` | `string` | No |  |
| `location_name` | `string` | No |  |
| `name` | `string` | No |  |
| `team` | `array` | No |  |
| `team_name` | `string` | No |  |
| `venue` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Team()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Team()->load(["id" => "team_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TeamEntity`

Create a new `TeamEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new NhlApiDocumentationSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

