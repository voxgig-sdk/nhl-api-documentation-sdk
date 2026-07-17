-- NhlApiDocumentation SDK exists test

local sdk = require("nhl-api-documentation_sdk")

describe("NhlApiDocumentationSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
