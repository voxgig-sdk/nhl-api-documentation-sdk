# Roster entity test

require "minitest/autorun"
require "json"
require_relative "../NhlApiDocumentation_sdk"
require_relative "runner"

class RosterEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NhlApiDocumentationSDK.test(nil, nil)
    ent = testsdk.Roster(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = roster_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "roster." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NHLAPIDOCUMENTATION_TEST_ROSTER_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    roster_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.roster")))
    roster_ref01_data = nil
    if roster_ref01_data_raw.length > 0
      roster_ref01_data = Helpers.to_map(roster_ref01_data_raw[0][1])
    end

    # LIST
    roster_ref01_ent = client.Roster(nil)
    roster_ref01_match = {
      "team_id" => setup[:idmap]["team01"],
    }

    roster_ref01_list_result, err = roster_ref01_ent.list(roster_ref01_match, nil)
    assert_nil err
    assert roster_ref01_list_result.is_a?(Array)

  end
end

def roster_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "roster", "RosterTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NhlApiDocumentationSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["roster01", "roster02", "roster03", "team01", "team02", "team03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["NHLAPIDOCUMENTATION_TEST_ROSTER_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NHLAPIDOCUMENTATION_TEST_ROSTER_ENTID" => idmap,
    "NHLAPIDOCUMENTATION_TEST_LIVE" => "FALSE",
    "NHLAPIDOCUMENTATION_TEST_EXPLAIN" => "FALSE",
    "NHLAPIDOCUMENTATION_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["NHLAPIDOCUMENTATION_TEST_ROSTER_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["NHLAPIDOCUMENTATION_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["NHLAPIDOCUMENTATION_APIKEY"],
      },
      extra || {},
    ])
    client = NhlApiDocumentationSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["NHLAPIDOCUMENTATION_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["NHLAPIDOCUMENTATION_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
