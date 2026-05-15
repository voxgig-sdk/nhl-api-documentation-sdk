<?php
declare(strict_types=1);

// NhlApiDocumentation SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class NhlApiDocumentationFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new NhlApiDocumentationBaseFeature();
            case "test":
                return new NhlApiDocumentationTestFeature();
            default:
                return new NhlApiDocumentationBaseFeature();
        }
    }
}
