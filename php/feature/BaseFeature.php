<?php
declare(strict_types=1);

// NhlApiDocumentation SDK base feature

class NhlApiDocumentationBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(NhlApiDocumentationContext $ctx, array $options): void {}
    public function PostConstruct(NhlApiDocumentationContext $ctx): void {}
    public function PostConstructEntity(NhlApiDocumentationContext $ctx): void {}
    public function SetData(NhlApiDocumentationContext $ctx): void {}
    public function GetData(NhlApiDocumentationContext $ctx): void {}
    public function GetMatch(NhlApiDocumentationContext $ctx): void {}
    public function SetMatch(NhlApiDocumentationContext $ctx): void {}
    public function PrePoint(NhlApiDocumentationContext $ctx): void {}
    public function PreSpec(NhlApiDocumentationContext $ctx): void {}
    public function PreRequest(NhlApiDocumentationContext $ctx): void {}
    public function PreResponse(NhlApiDocumentationContext $ctx): void {}
    public function PreResult(NhlApiDocumentationContext $ctx): void {}
    public function PreDone(NhlApiDocumentationContext $ctx): void {}
    public function PreUnexpected(NhlApiDocumentationContext $ctx): void {}
}
