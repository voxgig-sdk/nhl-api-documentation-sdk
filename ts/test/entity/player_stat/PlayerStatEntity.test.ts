
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { NhlApiDocumentationSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('PlayerStatEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NHLAPIDOCUMENTATION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NHLAPIDOCUMENTATION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NhlApiDocumentationSDK.test()
    const ent = testsdk.PlayerStat()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NHL_API_DOCUMENTATION_TEST_LIVE
    for (const op of ['list']) {
      if (maybeSkipControl(t, 'entityOp', 'player_stat.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set NHL_API_DOCUMENTATION_TEST_PLAYER_STAT_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let player_stat_ref01_data = Object.values(setup.data.existing.player_stat)[0] as any

    // LIST
    const player_stat_ref01_ent = client.PlayerStat()
    const player_stat_ref01_match: any = {}
    player_stat_ref01_match['person_id'] = setup.idmap['person01']

    const player_stat_ref01_list = await player_stat_ref01_ent.list(player_stat_ref01_match)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/player_stat/PlayerStatTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = NhlApiDocumentationSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['player_stat01','player_stat02','player_stat03','person01','person02','person03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['NHL_API_DOCUMENTATION_TEST_PLAYER_STAT_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'NHL_API_DOCUMENTATION_TEST_PLAYER_STAT_ENTID': idmap,
    'NHL_API_DOCUMENTATION_TEST_LIVE': 'FALSE',
    'NHL_API_DOCUMENTATION_TEST_EXPLAIN': 'FALSE',
    'NHL_API_DOCUMENTATION_APIKEY': 'NONE',
  })

  idmap = env['NHL_API_DOCUMENTATION_TEST_PLAYER_STAT_ENTID']

  const live = 'TRUE' === env.NHL_API_DOCUMENTATION_TEST_LIVE

  if (live) {
    client = new NhlApiDocumentationSDK(merge([
      {
        apikey: env.NHL_API_DOCUMENTATION_APIKEY,
      },
      extra
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.NHL_API_DOCUMENTATION_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
