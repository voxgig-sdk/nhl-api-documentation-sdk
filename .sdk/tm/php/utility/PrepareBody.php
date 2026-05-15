<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility: prepare_body

class NhlApiDocumentationPrepareBody
{
    public static function call(NhlApiDocumentationContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
