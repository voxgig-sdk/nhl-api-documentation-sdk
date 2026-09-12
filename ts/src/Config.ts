
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'NhlApiDocumentation',
        slug: "nhl-api-documentation",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      },
      "transport": "base"
    },

  }


  options = {
    base: "https://statsapi.web.nhl.com/api/v1",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      conference: {
      },

      division: {
      },

      game: {
      },

      player: {
      },

      player_stat: {
      },

      roster: {
      },

      schedule: {
      },

      standing: {
      },

      team: {
      },

    }
  }


  entity = {
    "conference": {
      "fields": [
        {
          "name": "conferences",
          "type": "`$ARRAY`"
        },
        {
          "name": "copyright",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$INTEGER`"
        },
        {
          "name": "link",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "conference",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/conferences",
              "segments": [
                {
                  "lit": "conferences"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.conferences`"
              },
              "parts": [
                "conferences"
              ]
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/conferences/{id}",
              "segments": [
                {
                  "lit": "conferences"
                },
                {
                  "var": "id"
                }
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "conferences",
                "{id}"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "division": {
      "fields": [
        {
          "name": "copyright",
          "type": "`$STRING`"
        },
        {
          "name": "divisions",
          "type": "`$ARRAY`"
        },
        {
          "name": "id",
          "type": "`$INTEGER`"
        },
        {
          "name": "link",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "division",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/divisions",
              "segments": [
                {
                  "lit": "divisions"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.divisions`"
              },
              "parts": [
                "divisions"
              ]
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/divisions/{id}",
              "segments": [
                {
                  "lit": "divisions"
                },
                {
                  "var": "id"
                }
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "divisions",
                "{id}"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "game": {
      "fields": [
        {
          "name": "away",
          "type": "`$OBJECT`"
        },
        {
          "name": "copyright",
          "type": "`$STRING`"
        },
        {
          "name": "gameData",
          "type": "`$OBJECT`"
        },
        {
          "name": "gamePk",
          "type": "`$INTEGER`"
        },
        {
          "name": "home",
          "type": "`$OBJECT`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "link",
          "type": "`$STRING`"
        },
        {
          "name": "liveData",
          "type": "`$OBJECT`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "game",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/game/{id}/boxscore",
              "segments": [
                {
                  "lit": "game"
                },
                {
                  "var": "id"
                },
                {
                  "lit": "boxscore"
                }
              ],
              "select": {
                "$action": "boxscore",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.teams`"
              },
              "parts": [
                "game",
                "{id}",
                "boxscore"
              ]
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/game/{id}/feed/live",
              "segments": [
                {
                  "lit": "game"
                },
                {
                  "var": "id"
                },
                {
                  "lit": "feed"
                },
                {
                  "lit": "live"
                }
              ],
              "select": {
                "$action": "feed_live",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "game",
                "{id}",
                "feed",
                "live"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "player": {
      "fields": [
        {
          "name": "copyright",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "people",
          "type": "`$ARRAY`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "player",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/people/{id}",
              "segments": [
                {
                  "lit": "people"
                },
                {
                  "var": "id"
                }
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "people",
                "{id}"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "player_stat": {
      "fields": [
        {
          "name": "splits",
          "type": "`$ARRAY`"
        },
        {
          "name": "type",
          "type": "`$OBJECT`"
        }
      ],
      "name": "player_stat",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "person_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "season",
                    "orig": "season",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "stat",
                    "orig": "stat",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/people/{id}/stats",
              "rename": {
                "param": {
                  "id": "person_id"
                }
              },
              "segments": [
                {
                  "lit": "people"
                },
                {
                  "var": "person_id"
                },
                {
                  "lit": "stats"
                }
              ],
              "select": {
                "exist": [
                  "person_id",
                  "season",
                  "stat"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.stats`"
              },
              "parts": [
                "people",
                "{person_id}",
                "stats"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "person"
          ]
        ]
      }
    },
    "roster": {
      "fields": [
        {
          "name": "jerseyNumber",
          "type": "`$STRING`"
        },
        {
          "name": "person",
          "type": "`$OBJECT`"
        },
        {
          "name": "position",
          "type": "`$OBJECT`"
        }
      ],
      "name": "roster",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "team_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "season",
                    "orig": "season",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/teams/{id}/roster",
              "rename": {
                "param": {
                  "id": "team_id"
                }
              },
              "segments": [
                {
                  "lit": "teams"
                },
                {
                  "var": "team_id"
                },
                {
                  "lit": "roster"
                }
              ],
              "select": {
                "exist": [
                  "season",
                  "team_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.roster`"
              },
              "parts": [
                "teams",
                "{team_id}",
                "roster"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "team"
          ]
        ]
      }
    },
    "schedule": {
      "fields": [
        {
          "format": "date",
          "name": "date",
          "type": "`$STRING`"
        },
        {
          "name": "games",
          "type": "`$ARRAY`"
        },
        {
          "name": "totalEvents",
          "type": "`$INTEGER`"
        },
        {
          "name": "totalGames",
          "type": "`$INTEGER`"
        },
        {
          "name": "totalItems",
          "type": "`$INTEGER`"
        },
        {
          "name": "totalMatches",
          "type": "`$INTEGER`"
        }
      ],
      "name": "schedule",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "end_date",
                    "orig": "end_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "season",
                    "orig": "season",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "start_date",
                    "orig": "start_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "team_id",
                    "orig": "team_id",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/schedule",
              "segments": [
                {
                  "lit": "schedule"
                }
              ],
              "select": {
                "exist": [
                  "end_date",
                  "season",
                  "start_date",
                  "team_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.dates`"
              },
              "parts": [
                "schedule"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "standing": {
      "fields": [
        {
          "name": "conference",
          "type": "`$OBJECT`"
        },
        {
          "name": "division",
          "type": "`$OBJECT`"
        },
        {
          "name": "teamRecords",
          "type": "`$ARRAY`"
        }
      ],
      "name": "standing",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "season",
                    "orig": "season",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/standings",
              "segments": [
                {
                  "lit": "standings"
                }
              ],
              "select": {
                "exist": [
                  "season"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.records`"
              },
              "parts": [
                "standings"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "team": {
      "fields": [
        {
          "name": "abbreviation",
          "type": "`$STRING`"
        },
        {
          "name": "conference",
          "type": "`$OBJECT`"
        },
        {
          "name": "copyright",
          "type": "`$STRING`"
        },
        {
          "name": "division",
          "type": "`$OBJECT`"
        },
        {
          "name": "firstYearOfPlay",
          "type": "`$STRING`"
        },
        {
          "name": "franchise",
          "type": "`$OBJECT`"
        },
        {
          "name": "id",
          "type": "`$INTEGER`"
        },
        {
          "name": "link",
          "type": "`$STRING`"
        },
        {
          "name": "locationName",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        },
        {
          "name": "teamName",
          "type": "`$STRING`"
        },
        {
          "name": "teams",
          "type": "`$ARRAY`"
        },
        {
          "name": "venue",
          "type": "`$OBJECT`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "team",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "expand",
                    "orig": "expand",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "season",
                    "orig": "season",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/teams",
              "segments": [
                {
                  "lit": "teams"
                }
              ],
              "select": {
                "exist": [
                  "expand",
                  "season"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.teams`"
              },
              "parts": [
                "teams"
              ]
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "expand",
                    "orig": "expand",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/teams/{id}",
              "segments": [
                {
                  "lit": "teams"
                },
                {
                  "var": "id"
                }
              ],
              "select": {
                "exist": [
                  "expand",
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "teams",
                "{id}"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

