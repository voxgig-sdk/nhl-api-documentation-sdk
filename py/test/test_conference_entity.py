# Conference entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from nhlapidocumentation_sdk import NhlApiDocumentationSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestConferenceEntity:

    def test_should_create_instance(self):
        testsdk = NhlApiDocumentationSDK.test(None, None)
        ent = testsdk.Conference(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _conference_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "conference." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set NHLAPIDOCUMENTATION_TEST_CONFERENCE_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        conference_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.conference")))
        conference_ref01_data = None
        if len(conference_ref01_data_raw) > 0:
            conference_ref01_data = helpers.to_map(conference_ref01_data_raw[0][1])

        # LIST
        conference_ref01_ent = client.Conference(None)
        conference_ref01_match = {}

        conference_ref01_list_result, err = conference_ref01_ent.list(conference_ref01_match, None)
        assert err is None
        assert isinstance(conference_ref01_list_result, list)

        # LOAD
        conference_ref01_match_dt0 = {
            "id": conference_ref01_data["id"],
        }
        conference_ref01_data_dt0_loaded, err = conference_ref01_ent.load(conference_ref01_match_dt0, None)
        assert err is None
        conference_ref01_data_dt0_load_result = helpers.to_map(conference_ref01_data_dt0_loaded)
        assert conference_ref01_data_dt0_load_result is not None
        assert conference_ref01_data_dt0_load_result["id"] == conference_ref01_data["id"]



def _conference_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/conference/ConferenceTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = NhlApiDocumentationSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["conference01", "conference02", "conference03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "NHLAPIDOCUMENTATION_TEST_CONFERENCE_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "NHLAPIDOCUMENTATION_TEST_CONFERENCE_ENTID": idmap,
        "NHLAPIDOCUMENTATION_TEST_LIVE": "FALSE",
        "NHLAPIDOCUMENTATION_TEST_EXPLAIN": "FALSE",
        "NHLAPIDOCUMENTATION_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("NHLAPIDOCUMENTATION_TEST_CONFERENCE_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("NHLAPIDOCUMENTATION_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
                "apikey": env.get("NHLAPIDOCUMENTATION_APIKEY"),
            },
            extra or {},
        ])
        client = NhlApiDocumentationSDK(helpers.to_map(merged_opts))

    _live = env.get("NHLAPIDOCUMENTATION_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("NHLAPIDOCUMENTATION_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
