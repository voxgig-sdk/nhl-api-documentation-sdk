# NhlApiDocumentation PHP SDK Reference

Complete API reference for the NhlApiDocumentation PHP SDK.


## NhlApiDocumentationSDK

### Constructor

```php
require_once __DIR__ . '/nhl-api-documentation_sdk.php';

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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

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
$conference = $client->conference();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | ``$ARRAY`` | No |  |
| `copyright` | ``$STRING`` | No |  |
| `id` | ``$INTEGER`` | No |  |
| `link` | ``$STRING`` | No |  |
| `name` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->conference()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->conference()->load(["id" => "conference_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ConferenceEntity`

Create a new `ConferenceEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## DivisionEntity

```php
$division = $client->division();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | ``$STRING`` | No |  |
| `division` | ``$ARRAY`` | No |  |
| `id` | ``$INTEGER`` | No |  |
| `link` | ``$STRING`` | No |  |
| `name` | ``$STRING`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->division()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->division()->load(["id" => "division_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): DivisionEntity`

Create a new `DivisionEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## GameEntity

```php
$game = $client->game();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | ``$STRING`` | No |  |
| `game_data` | ``$OBJECT`` | No |  |
| `game_pk` | ``$INTEGER`` | No |  |
| `link` | ``$STRING`` | No |  |
| `live_data` | ``$OBJECT`` | No |  |
| `team` | ``$OBJECT`` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->game()->load(["id" => "game_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): GameEntity`

Create a new `GameEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PlayerEntity

```php
$player = $client->player();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | ``$STRING`` | No |  |
| `person` | ``$ARRAY`` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->player()->load(["id" => "player_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PlayerEntity`

Create a new `PlayerEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## PlayerStatEntity

```php
$player_stat = $client->player_stat();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | ``$ARRAY`` | No |  |
| `type` | ``$OBJECT`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->player_stat()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): PlayerStatEntity`

Create a new `PlayerStatEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## RosterEntity

```php
$roster = $client->roster();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | ``$STRING`` | No |  |
| `person` | ``$OBJECT`` | No |  |
| `position` | ``$OBJECT`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->roster()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): RosterEntity`

Create a new `RosterEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## ScheduleEntity

```php
$schedule = $client->schedule();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | ``$STRING`` | No |  |
| `game` | ``$ARRAY`` | No |  |
| `total_event` | ``$INTEGER`` | No |  |
| `total_game` | ``$INTEGER`` | No |  |
| `total_item` | ``$INTEGER`` | No |  |
| `total_match` | ``$INTEGER`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->schedule()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): ScheduleEntity`

Create a new `ScheduleEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## StandingEntity

```php
$standing = $client->standing();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | ``$OBJECT`` | No |  |
| `division` | ``$OBJECT`` | No |  |
| `team_record` | ``$ARRAY`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->standing()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): StandingEntity`

Create a new `StandingEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## TeamEntity

```php
$team = $client->team();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | ``$STRING`` | No |  |
| `conference` | ``$OBJECT`` | No |  |
| `copyright` | ``$STRING`` | No |  |
| `division` | ``$OBJECT`` | No |  |
| `first_year_of_play` | ``$STRING`` | No |  |
| `franchise` | ``$OBJECT`` | No |  |
| `id` | ``$INTEGER`` | No |  |
| `link` | ``$STRING`` | No |  |
| `location_name` | ``$STRING`` | No |  |
| `name` | ``$STRING`` | No |  |
| `team` | ``$ARRAY`` | No |  |
| `team_name` | ``$STRING`` | No |  |
| `venue` | ``$OBJECT`` | No |  |

### Operations

#### `list(array $reqmatch, ?array $ctrl = null): mixed`

List entities matching the given criteria. Returns an array. Throws on error.

```php
$results = $client->team()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->team()->load(["id" => "team_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): TeamEntity`

Create a new `TeamEntity` instance with the same client and
options.

#### `getName(): string`

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

