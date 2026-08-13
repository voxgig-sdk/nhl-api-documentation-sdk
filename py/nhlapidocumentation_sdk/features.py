# NhlApiDocumentation SDK feature factory

from nhlapidocumentation_sdk.feature.base_feature import NhlApiDocumentationBaseFeature
from nhlapidocumentation_sdk.feature.test_feature import NhlApiDocumentationTestFeature


def _make_feature(name):
    features = {
        "base": lambda: NhlApiDocumentationBaseFeature(),
        "test": lambda: NhlApiDocumentationTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
