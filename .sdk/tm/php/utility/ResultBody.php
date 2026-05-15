<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility: result_body

class NhlApiDocumentationResultBody
{
    public static function call(NhlApiDocumentationContext $ctx): ?NhlApiDocumentationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
