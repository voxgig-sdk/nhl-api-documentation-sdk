package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "NhlApiDocumentation",
			"slug": "nhl-api-documentation",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://statsapi.web.nhl.com/api/v1",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"conference": map[string]any{},
				"division": map[string]any{},
				"game": map[string]any{},
				"player": map[string]any{},
				"player_stat": map[string]any{},
				"roster": map[string]any{},
				"schedule": map[string]any{},
				"standing": map[string]any{},
				"team": map[string]any{},
			},
		},
		"entity": map[string]any{
			"conference": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "conferences",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "copyright",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "link",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
				},
				"name": "conference",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/conferences",
								"parts": []any{
									"conferences",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.conferences`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/conferences/{id}",
								"parts": []any{
									"conferences",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"division": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "copyright",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "divisions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "link",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
				},
				"name": "division",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/divisions",
								"parts": []any{
									"divisions",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.divisions`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/divisions/{id}",
								"parts": []any{
									"divisions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"game": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "away",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "copyright",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gameData",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gamePk",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "home",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "link",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "liveData",
						"type": "`$OBJECT`",
					},
				},
				"name": "game",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/game/{id}/boxscore",
								"parts": []any{
									"game",
									"{id}",
									"boxscore",
								},
								"select": map[string]any{
									"$action": "boxscore",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.teams`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/game/{id}/feed/live",
								"parts": []any{
									"game",
									"{id}",
									"feed",
									"live",
								},
								"select": map[string]any{
									"$action": "feed_live",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"player": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "copyright",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "people",
						"type": "`$ARRAY`",
					},
				},
				"name": "player",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/people/{id}",
								"parts": []any{
									"people",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"player_stat": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "splits",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "type",
						"type": "`$OBJECT`",
					},
				},
				"name": "player_stat",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "person_id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "season",
											"orig": "season",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "stat",
											"orig": "stat",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/people/{id}/stats",
								"parts": []any{
									"people",
									"{person_id}",
									"stats",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "person_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"person_id",
										"season",
										"stat",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.stats`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"person",
						},
					},
				},
			},
			"roster": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "jerseyNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "person",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "position",
						"type": "`$OBJECT`",
					},
				},
				"name": "roster",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "team_id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "season",
											"orig": "season",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/teams/{id}/roster",
								"parts": []any{
									"teams",
									"{team_id}",
									"roster",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "team_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"season",
										"team_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.roster`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"team",
						},
					},
				},
			},
			"schedule": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "games",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "totalEvents",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "totalGames",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "totalItems",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "totalMatches",
						"type": "`$INTEGER`",
					},
				},
				"name": "schedule",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "end_date",
											"orig": "end_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "season",
											"orig": "season",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "start_date",
											"orig": "start_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "team_id",
											"orig": "team_id",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/schedule",
								"parts": []any{
									"schedule",
								},
								"select": map[string]any{
									"exist": []any{
										"end_date",
										"season",
										"start_date",
										"team_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.dates`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"standing": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "conference",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "division",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "teamRecords",
						"type": "`$ARRAY`",
					},
				},
				"name": "standing",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "season",
											"orig": "season",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/standings",
								"parts": []any{
									"standings",
								},
								"select": map[string]any{
									"exist": []any{
										"season",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.records`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"team": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "abbreviation",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "conference",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "copyright",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "division",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "firstYearOfPlay",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "franchise",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "link",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "locationName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "teamName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "teams",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "venue",
						"type": "`$OBJECT`",
					},
				},
				"name": "team",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "expand",
											"orig": "expand",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "season",
											"orig": "season",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/teams",
								"parts": []any{
									"teams",
								},
								"select": map[string]any{
									"exist": []any{
										"expand",
										"season",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.teams`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "expand",
											"orig": "expand",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/teams/{id}",
								"parts": []any{
									"teams",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"expand",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
