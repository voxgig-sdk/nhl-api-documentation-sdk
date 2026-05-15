# NhlApiDocumentation SDK feature factory

from feature.base_feature import NhlApiDocumentationBaseFeature
from feature.test_feature import NhlApiDocumentationTestFeature


def _make_feature(name):
    features = {
        "base": lambda: NhlApiDocumentationBaseFeature(),
        "test": lambda: NhlApiDocumentationTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
