package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/nhl-api-documentation-sdk/go"
	"github.com/voxgig-sdk/nhl-api-documentation-sdk/go/core"

	vs "github.com/voxgig-sdk/nhl-api-documentation-sdk/go/utility/struct"
)

func TestDivisionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Division(nil)
		if ent == nil {
			t.Fatal("expected non-nil DivisionEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := divisionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "division." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		divisionRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.division", setup.data)))
		var divisionRef01Data map[string]any
		if len(divisionRef01DataRaw) > 0 {
			divisionRef01Data = core.ToMapAny(divisionRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = divisionRef01Data

		// LIST
		divisionRef01Ent := client.Division(nil)
		divisionRef01Match := map[string]any{}

		divisionRef01ListResult, err := divisionRef01Ent.List(divisionRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, divisionRef01ListOk := divisionRef01ListResult.([]any)
		if !divisionRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", divisionRef01ListResult)
		}

		// LOAD
		divisionRef01MatchDt0 := map[string]any{
			"id": divisionRef01Data["id"],
		}
		divisionRef01DataDt0Loaded, err := divisionRef01Ent.Load(divisionRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		divisionRef01DataDt0LoadResult := core.ToMapAny(divisionRef01DataDt0Loaded)
		if divisionRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if divisionRef01DataDt0LoadResult["id"] != divisionRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func divisionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "division", "DivisionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read division test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse division test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"division01", "division02", "division03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID": idmap,
		"NHLAPIDOCUMENTATION_TEST_LIVE":      "FALSE",
		"NHLAPIDOCUMENTATION_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["NHLAPIDOCUMENTATION_TEST_DIVISION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["NHLAPIDOCUMENTATION_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
			},
			extra,
		})
		client = sdk.NewNhlApiDocumentationSDK(core.ToMapAny(mergedOpts))
	}

	live := env["NHLAPIDOCUMENTATION_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["NHLAPIDOCUMENTATION_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
