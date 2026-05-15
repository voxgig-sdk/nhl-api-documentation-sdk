# ProjectName SDK exists test

import pytest
from nhlapidocumentation_sdk import NhlApiDocumentationSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = NhlApiDocumentationSDK.test(None, None)
        assert testsdk is not None
