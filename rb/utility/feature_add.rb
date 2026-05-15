# NhlApiDocumentation SDK utility: feature_add
module NhlApiDocumentationUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
