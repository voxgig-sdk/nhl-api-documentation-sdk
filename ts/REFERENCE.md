# NhlApiDocumentation TypeScript SDK Reference

Complete API reference for the NhlApiDocumentation TypeScript SDK.


## NhlApiDocumentationSDK

### Constructor

```ts
new NhlApiDocumentationSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NhlApiDocumentationSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NhlApiDocumentationSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NhlApiDocumentationSDK` instance in test mode.


### Instance Methods

#### `Conference(data?: object)`

Create a new `Conference` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ConferenceEntity` instance.

#### `Division(data?: object)`

Create a new `Division` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DivisionEntity` instance.

#### `Game(data?: object)`

Create a new `Game` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GameEntity` instance.

#### `Player(data?: object)`

Create a new `Player` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlayerEntity` instance.

#### `PlayerStat(data?: object)`

Create a new `PlayerStat` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlayerStatEntity` instance.

#### `Roster(data?: object)`

Create a new `Roster` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RosterEntity` instance.

#### `Schedule(data?: object)`

Create a new `Schedule` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ScheduleEntity` instance.

#### `Standing(data?: object)`

Create a new `Standing` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `StandingEntity` instance.

#### `Team(data?: object)`

Create a new `Team` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TeamEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NhlApiDocumentationSDK.test()`.

**Returns:** `NhlApiDocumentationSDK` instance in test mode.


---

## ConferenceEntity

```ts
const conference = client.Conference()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `any[]` | No |  |
| `copyright` | `string` | No |  |
| `id` | `number` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Conference().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Conference().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ConferenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DivisionEntity

```ts
const division = client.Division()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `division` | `any[]` | No |  |
| `id` | `number` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Division().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Division().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DivisionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GameEntity

```ts
const game = client.Game()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `game_data` | `Record<string, any>` | No |  |
| `game_pk` | `number` | No |  |
| `link` | `string` | No |  |
| `live_data` | `Record<string, any>` | No |  |
| `team` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Game().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GameEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlayerEntity

```ts
const player = client.Player()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `copyright` | `string` | No |  |
| `person` | `any[]` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Player().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlayerEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlayerStatEntity

```ts
const player_stat = client.PlayerStat()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `split` | `any[]` | No |  |
| `type` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PlayerStat().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlayerStatEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RosterEntity

```ts
const roster = client.Roster()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jersey_number` | `string` | No |  |
| `person` | `Record<string, any>` | No |  |
| `position` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Roster().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RosterEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ScheduleEntity

```ts
const schedule = client.Schedule()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `date` | `string` | No |  |
| `game` | `any[]` | No |  |
| `total_event` | `number` | No |  |
| `total_game` | `number` | No |  |
| `total_item` | `number` | No |  |
| `total_match` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Schedule().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ScheduleEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## StandingEntity

```ts
const standing = client.Standing()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `conference` | `Record<string, any>` | No |  |
| `division` | `Record<string, any>` | No |  |
| `team_record` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Standing().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `StandingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TeamEntity

```ts
const team = client.Team()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No |  |
| `conference` | `Record<string, any>` | No |  |
| `copyright` | `string` | No |  |
| `division` | `Record<string, any>` | No |  |
| `first_year_of_play` | `string` | No |  |
| `franchise` | `Record<string, any>` | No |  |
| `id` | `number` | No |  |
| `link` | `string` | No |  |
| `location_name` | `string` | No |  |
| `name` | `string` | No |  |
| `team` | `any[]` | No |  |
| `team_name` | `string` | No |  |
| `venue` | `Record<string, any>` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Team().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Team().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TeamEntity` instance with the same client and
options.

#### `client()`

Return the parent `NhlApiDocumentationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NhlApiDocumentationSDK({
  feature: {
    test: { active: true },
  }
})
```

