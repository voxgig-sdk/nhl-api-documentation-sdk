<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility: result_headers

class NhlApiDocumentationResultHeaders
{
    public static function call(NhlApiDocumentationContext $ctx): ?NhlApiDocumentationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
