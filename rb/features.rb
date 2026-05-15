# NhlApiDocumentation SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module NhlApiDocumentationFeatures
  def self.make_feature(name)
    case name
    when "base"
      NhlApiDocumentationBaseFeature.new
    when "test"
      NhlApiDocumentationTestFeature.new
    else
      NhlApiDocumentationBaseFeature.new
    end
  end
end
