# NhlApiDocumentation SDK utility: make_context
require_relative '../core/context'
module NhlApiDocumentationUtilities
  MakeContext = ->(ctxmap, basectx) {
    NhlApiDocumentationContext.new(ctxmap, basectx)
  }
end
