<?php
declare(strict_types=1);

// NhlApiDocumentation SDK exists test

require_once __DIR__ . '/../nhlapidocumentation_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = NhlApiDocumentationSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
