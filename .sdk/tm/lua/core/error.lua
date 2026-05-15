-- NhlApiDocumentation SDK error

local NhlApiDocumentationError = {}
NhlApiDocumentationError.__index = NhlApiDocumentationError


function NhlApiDocumentationError.new(code, msg, ctx)
  local self = setmetatable({}, NhlApiDocumentationError)
  self.is_sdk_error = true
  self.sdk = "NhlApiDocumentation"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function NhlApiDocumentationError:error()
  return self.msg
end


function NhlApiDocumentationError:__tostring()
  return self.msg
end


return NhlApiDocumentationError
