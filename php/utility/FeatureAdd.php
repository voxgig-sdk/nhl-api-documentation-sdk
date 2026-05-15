<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility: feature_add

class NhlApiDocumentationFeatureAdd
{
    public static function call(NhlApiDocumentationContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
