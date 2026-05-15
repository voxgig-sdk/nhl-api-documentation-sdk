# NhlApiDocumentation SDK exists test

require "minitest/autorun"
require_relative "../NhlApiDocumentation_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = NhlApiDocumentationSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
