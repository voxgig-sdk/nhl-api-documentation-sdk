<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class NhlApiDocumentationMakeContext
{
    public static function call(array $ctxmap, ?NhlApiDocumentationContext $basectx): NhlApiDocumentationContext
    {
        return new NhlApiDocumentationContext($ctxmap, $basectx);
    }
}
