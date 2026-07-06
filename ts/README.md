# NhlApiDocumentation TypeScript SDK



The TypeScript SDK for the NhlApiDocumentation API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Conference()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases](https://github.com/voxgig-sdk/nhl-api-documentation-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { NhlApiDocumentationSDK } from '@voxgig-sdk/nhl-api-documentation'

const client = new NhlApiDocumentationSDK()
```

### 2. List conference records

`list()` resolves to an array of Conference objects — iterate it directly:

```ts
const conferences = await client.Conference().list()

for (const conference of conferences) {
  console.log(conference)
}
```

### 3. Load a conference

`load()` returns the entity directly and throws on failure:

```ts
try {
  const conference = await client.Conference().load({ id: 1 })
  console.log(conference)
} catch (err) {
  console.error('load failed:', err)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const conferences = await client.Conference().list()
  console.log(conferences)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = NhlApiDocumentationSDK.test()

const conference = await client.Conference().list()
// conference is a bare entity populated with mock response data
console.log(conference)
```

You can also use the instance method:

```ts
const client = new NhlApiDocumentationSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Conference()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new NhlApiDocumentationSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NHL_API_DOCUMENTATION_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### NhlApiDocumentationSDK

#### Constructor

```ts
new NhlApiDocumentationSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Conference(data?)` | `ConferenceEntity` | Create a Conference entity instance. |
| `Division(data?)` | `DivisionEntity` | Create a Division entity instance. |
| `Game(data?)` | `GameEntity` | Create a Game entity instance. |
| `Player(data?)` | `PlayerEntity` | Create a Player entity instance. |
| `PlayerStat(data?)` | `PlayerStatEntity` | Create a PlayerStat entity instance. |
| `Roster(data?)` | `RosterEntity` | Create a Roster entity instance. |
| `Schedule(data?)` | `ScheduleEntity` | Create a Schedule entity instance. |
| `Standing(data?)` | `StandingEntity` | Create a Standing entity instance. |
| `Team(data?)` | `TeamEntity` | Create a Team entity instance. |
| `tester(testopts?, sdkopts?)` | `NhlApiDocumentationSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `NhlApiDocumentationSDK.test(testopts?, sdkopts?)` | `NhlApiDocumentationSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): NhlApiDocumentationSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Conference

| Field | Description |
| --- | --- |
| `conference` |  |
| `copyright` |  |
| `id` |  |
| `link` |  |
| `name` |  |

Operations: list, load.

API path: `/conferences`

#### Division

| Field | Description |
| --- | --- |
| `copyright` |  |
| `division` |  |
| `id` |  |
| `link` |  |
| `name` |  |

Operations: list, load.

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

Operations: load.

API path: `/game/{id}/boxscore`

#### Player

| Field | Description |
| --- | --- |
| `copyright` |  |
| `person` |  |

Operations: load.

API path: `/people/{id}`

#### PlayerStat

| Field | Description |
| --- | --- |
| `split` |  |
| `type` |  |

Operations: list.

API path: `/people/{id}/stats`

#### Roster

| Field | Description |
| --- | --- |
| `jersey_number` |  |
| `person` |  |
| `position` |  |

Operations: list.

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

Operations: list.

API path: `/schedule`

#### Standing

| Field | Description |
| --- | --- |
| `conference` |  |
| `division` |  |
| `team_record` |  |

Operations: list.

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

Operations: list, load.

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
| `conference` | `any[]` |  |
| `copyright` | `string` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `name` | `string` |  |

#### Example: Load

```ts
const conference = await client.Conference().load({ id: 1 })
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
| `copyright` | `string` |  |
| `division` | `any[]` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `name` | `string` |  |

#### Example: Load

```ts
const division = await client.Division().load({ id: 1 })
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
| `copyright` | `string` |  |
| `game_data` | `Record<string, any>` |  |
| `game_pk` | `number` |  |
| `link` | `string` |  |
| `live_data` | `Record<string, any>` |  |
| `team` | `Record<string, any>` |  |

#### Example: Load

```ts
const game = await client.Game().load({ id: 1 })
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
| `copyright` | `string` |  |
| `person` | `any[]` |  |

#### Example: Load

```ts
const player = await client.Player().load({ id: 1 })
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
| `split` | `any[]` |  |
| `type` | `Record<string, any>` |  |

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
| `jersey_number` | `string` |  |
| `person` | `Record<string, any>` |  |
| `position` | `Record<string, any>` |  |

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
| `date` | `string` |  |
| `game` | `any[]` |  |
| `total_event` | `number` |  |
| `total_game` | `number` |  |
| `total_item` | `number` |  |
| `total_match` | `number` |  |

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
| `conference` | `Record<string, any>` |  |
| `division` | `Record<string, any>` |  |
| `team_record` | `any[]` |  |

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
| `abbreviation` | `string` |  |
| `conference` | `Record<string, any>` |  |
| `copyright` | `string` |  |
| `division` | `Record<string, any>` |  |
| `first_year_of_play` | `string` |  |
| `franchise` | `Record<string, any>` |  |
| `id` | `number` |  |
| `link` | `string` |  |
| `location_name` | `string` |  |
| `name` | `string` |  |
| `team` | `any[]` |  |
| `team_name` | `string` |  |
| `venue` | `Record<string, any>` |  |

#### Example: Load

```ts
const team = await client.Team().load({ id: 1 })
```

#### Example: List

```ts
const teams = await client.Team().list()
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
nhl-api-documentation/
├── src/
│   ├── NhlApiDocumentationSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { NhlApiDocumentationSDK } from '@voxgig-sdk/nhl-api-documentation'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const conference = client.Conference()
await conference.list()

// conference.data() now returns the conference data from the last `list`
// conference.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
