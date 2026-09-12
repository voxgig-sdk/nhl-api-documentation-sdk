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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "conferences",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.conferences`",
								},
								"parts": []any{
									"conferences",
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
								"segments": []any{
									map[string]any{
										"lit": "conferences",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"conferences",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "divisions",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.divisions`",
								},
								"parts": []any{
									"divisions",
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
								"segments": []any{
									map[string]any{
										"lit": "divisions",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"divisions",
									"{id}",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "game",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "boxscore",
									},
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
								"parts": []any{
									"game",
									"{id}",
									"boxscore",
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
								"segments": []any{
									map[string]any{
										"lit": "game",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "feed",
									},
									map[string]any{
										"lit": "live",
									},
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
								"parts": []any{
									"game",
									"{id}",
									"feed",
									"live",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "people",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"people",
									"{id}",
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
								"rename": map[string]any{
									"param": map[string]any{
										"id": "person_id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "people",
									},
									map[string]any{
										"var": "person_id",
									},
									map[string]any{
										"lit": "stats",
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
								"parts": []any{
									"people",
									"{person_id}",
									"stats",
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
								"rename": map[string]any{
									"param": map[string]any{
										"id": "team_id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "teams",
									},
									map[string]any{
										"var": "team_id",
									},
									map[string]any{
										"lit": "roster",
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
								"parts": []any{
									"teams",
									"{team_id}",
									"roster",
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
						"format": "date",
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
								"segments": []any{
									map[string]any{
										"lit": "schedule",
									},
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
								"parts": []any{
									"schedule",
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
								"segments": []any{
									map[string]any{
										"lit": "standings",
									},
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
								"parts": []any{
									"standings",
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
				"id": map[string]any{
					"field": "id",
					"name": "id",
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
								"segments": []any{
									map[string]any{
										"lit": "teams",
									},
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
								"parts": []any{
									"teams",
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
								"segments": []any{
									map[string]any{
										"lit": "teams",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"teams",
									"{id}",
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

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
