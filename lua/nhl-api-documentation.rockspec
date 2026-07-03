package = "voxgig-sdk-nhl-api-documentation"
version = "0.0.1-1"
source = {
  -- git+https (GitHub dropped git:// in 2022); pin the install to the release
  -- tag pushed by `make publish`, and point at the lua/ subdir of the monorepo.
  url = "git+https://github.com/voxgig-sdk/nhl-api-documentation-sdk.git",
  tag = "lua/v0.0.1",
  dir = "nhl-api-documentation-sdk/lua"
}
description = {
  summary = "Unofficial generated Lua SDK for the NHL API Documentation public API. Not affiliated with or endorsed by the upstream API provider.",
  homepage = "https://github.com/voxgig-sdk/nhl-api-documentation-sdk",
  issues_url = "https://github.com/voxgig-sdk/nhl-api-documentation-sdk/issues",
  license = "MIT",
  labels = { "voxgig", "sdk", "generated-sdk", "openapi", "api-client", "nhl-api-documentation" }
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["nhl-api-documentation_sdk"] = "nhl-api-documentation_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
