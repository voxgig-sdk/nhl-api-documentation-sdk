package = "voxgig-sdk-nhl-api-documentation"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/nhl-api-documentation-sdk.git"
}
description = {
  summary = "NhlApiDocumentation SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
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
