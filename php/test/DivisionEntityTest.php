<?php
declare(strict_types=1);

// Division entity test

require_once __DIR__ . '/../nhlapidocumentation_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class DivisionEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NhlApiDocumentationSDK::test(null, null);
        $ent = $testsdk->Division(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = division_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "division." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $division_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.division")));
        $division_ref01_data = null;
        if (count($division_ref01_data_raw) > 0) {
            $division_ref01_data = Helpers::to_map($division_ref01_data_raw[0][1]);
        }

        // LIST
        $division_ref01_ent = $client->Division(null);
        $division_ref01_match = [];

        [$division_ref01_list_result, $err] = $division_ref01_ent->list($division_ref01_match, null);
        $this->assertNull($err);
        $this->assertIsArray($division_ref01_list_result);

        // LOAD
        $division_ref01_match_dt0 = [
            "id" => $division_ref01_data["id"],
        ];
        [$division_ref01_data_dt0_loaded, $err] = $division_ref01_ent->load($division_ref01_match_dt0, null);
        $this->assertNull($err);
        $division_ref01_data_dt0_load_result = Helpers::to_map($division_ref01_data_dt0_loaded);
        $this->assertNotNull($division_ref01_data_dt0_load_result);
        $this->assertEquals($division_ref01_data_dt0_load_result["id"], $division_ref01_data["id"]);

    }
}

function division_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/division/DivisionTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NhlApiDocumentationSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["division01", "division02", "division03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID" => $idmap,
        "NHLAPIDOCUMENTATION_TEST_LIVE" => "FALSE",
        "NHLAPIDOCUMENTATION_TEST_EXPLAIN" => "FALSE",
        "NHLAPIDOCUMENTATION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["NHLAPIDOCUMENTATION_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["NHLAPIDOCUMENTATION_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new NhlApiDocumentationSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["NHLAPIDOCUMENTATION_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["NHLAPIDOCUMENTATION_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
