# NhlApiDocumentation SDK utility: make_context

from core.context import NhlApiDocumentationContext


def make_context_util(ctxmap, basectx):
    return NhlApiDocumentationContext(ctxmap, basectx)
