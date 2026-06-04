# NhlApiDocumentation SDK

Query NHL teams, players, rosters, schedules, standings, and live game data without authentication

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About NHL API Documentation

The NHL API is a publicly accessible, unauthenticated set of HTTP endpoints used by the NHL's own apps and websites. There is no official developer programme; the most widely-used reference is the community-maintained [dword4/nhlapi](https://gitlab.com/dword4/nhlapi) project by Drew Hynes, which catalogues the routes and response shapes.

Note that the historical host `statsapi.web.nhl.com/api/v1` (the server configured for this SDK) has been deprecated by the NHL. The current host is `https://api-web.nhle.com/v1/` and existing clients should be migrated to it.

What you can pull from the API:

- Team metadata, conferences, and divisions
- Player biographies, current and historical rosters
- League schedule and individual game data, including live play-by-play
- Standings by season, conference, and division
- Aggregated player statistics

Endpoints are read-only GET requests returning JSON and do not require an API key. Because the API is undocumented officially, response shapes can change without notice; pinning to known-good fields and handling missing values defensively is recommended.

## Try it

**TypeScript**
```bash
npm install nhl-api-documentation
```

**Python**
```bash
pip install nhl-api-documentation-sdk
```

**PHP**
```bash
composer require voxgig/nhl-api-documentation-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/nhl-api-documentation-sdk/go
```

**Ruby**
```bash
gem install nhl-api-documentation-sdk
```

**Lua**
```bash
luarocks install nhl-api-documentation-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { NhlApiDocumentationSDK } from 'nhl-api-documentation'

const client = new NhlApiDocumentationSDK({})

// List all conferences
const conferences = await client.Conference().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o nhl-api-documentation-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "nhl-api-documentation": {
      "command": "/abs/path/to/nhl-api-documentation-mcp"
    }
  }
}
```

## Entities

The API exposes 9 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Conference** | An NHL conference (Eastern, Western) used to group divisions and teams in standings and team listings. | `/conferences` |
| **Division** | A division within a conference (e.g. Atlantic, Metropolitan, Central, Pacific) used to organise teams and standings. | `/divisions` |
| **Game** | A single NHL game resource exposing schedule entry, boxscore, and live play-by-play data. | `/game/{id}/boxscore` |
| **Player** | An individual NHL player resource with biographical details and career references. | `/people/{id}` |
| **PlayerStat** | Aggregated statistics for a player, typically by season and game type (regular season, playoffs). | `/people/{id}/stats` |
| **Roster** | The set of players currently or historically assigned to a team, usually scoped to a season. | `/teams/{id}/roster` |
| **Schedule** | League or team schedule listings, filterable by date range and used to discover game IDs. | `/schedule` |
| **Standing** | Team standings for a season, broken down by league, conference, division, or wildcard. | `/standings` |
| **Team** | An NHL franchise resource with metadata such as name, abbreviation, venue, division, and conference. | `/teams` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from nhlapidocumentation_sdk import NhlApiDocumentationSDK

client = NhlApiDocumentationSDK({})

# List all conferences
conferences, err = client.Conference(None).list(None, None)

# Load a specific conference
conference, err = client.Conference(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'nhlapidocumentation_sdk.php';

$client = new NhlApiDocumentationSDK([]);

// List all conferences
[$conferences, $err] = $client->Conference(null)->list(null, null);

// Load a specific conference
[$conference, $err] = $client->Conference(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/nhl-api-documentation-sdk/go"

client := sdk.NewNhlApiDocumentationSDK(map[string]any{})

// List all conferences
conferences, err := client.Conference(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "NhlApiDocumentation_sdk"

client = NhlApiDocumentationSDK.new({})

# List all conferences
conferences, err = client.Conference(nil).list(nil, nil)

# Load a specific conference
conference, err = client.Conference(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("nhl-api-documentation_sdk")

local client = sdk.new({})

-- List all conferences
local conferences, err = client:Conference(nil):list(nil, nil)

-- Load a specific conference
local conference, err = client:Conference(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = NhlApiDocumentationSDK.test()
const result = await client.Conference().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = NhlApiDocumentationSDK.test(None, None)
result, err = client.Conference(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = NhlApiDocumentationSDK::test(null, null);
[$result, $err] = $client->Conference(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Conference(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = NhlApiDocumentationSDK.test(nil, nil)
result, err = client.Conference(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Conference(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the NHL API Documentation

- Upstream: [https://gitlab.com/dword4/nhlapi](https://gitlab.com/dword4/nhlapi)

- The community documentation project ([dword4/nhlapi](https://gitlab.com/dword4/nhlapi)) is published under the MIT License.
- The underlying data is served by the NHL's public web endpoints; this SDK is unofficial and not affiliated with the NHL.
- Respect the NHL's terms of use when redistributing data and consider caching responses to avoid excessive load.

---

Generated from the NHL API Documentation OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
