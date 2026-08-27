<?php
declare(strict_types=1);

// NhlApiDocumentation SDK configuration

class NhlApiDocumentationConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "NhlApiDocumentation",
                "slug" => "nhl-api-documentation",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
          'transport' => 'base',
        ],
            ],
            "options" => [
                "base" => "https://statsapi.web.nhl.com/api/v1",
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "conference" => [],
                    "division" => [],
                    "game" => [],
                    "player" => [],
                    "player_stat" => [],
                    "roster" => [],
                    "schedule" => [],
                    "standing" => [],
                    "team" => [],
                ],
            ],
            "entity" => [
        'conference' => [
          'fields' => [
            [
              'name' => 'conferences',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'copyright',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'link',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'conference',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/conferences',
                  'parts' => [
                    'conferences',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.conferences`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/conferences/{id}',
                  'parts' => [
                    'conferences',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'division' => [
          'fields' => [
            [
              'name' => 'copyright',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'divisions',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'id',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'link',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'division',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/divisions',
                  'parts' => [
                    'divisions',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.divisions`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/divisions/{id}',
                  'parts' => [
                    'divisions',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'game' => [
          'fields' => [
            [
              'name' => 'away',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'copyright',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'gameData',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'gamePk',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'home',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'link',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'liveData',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'game',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/game/{id}/boxscore',
                  'parts' => [
                    'game',
                    '{id}',
                    'boxscore',
                  ],
                  'select' => [
                    '$action' => 'boxscore',
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.teams`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/game/{id}/feed/live',
                  'parts' => [
                    'game',
                    '{id}',
                    'feed',
                    'live',
                  ],
                  'select' => [
                    '$action' => 'feed_live',
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'player' => [
          'fields' => [
            [
              'name' => 'copyright',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'people',
              'type' => '`$ARRAY`',
            ],
          ],
          'name' => 'player',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/people/{id}',
                  'parts' => [
                    'people',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'player_stat' => [
          'fields' => [
            [
              'name' => 'splits',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'type',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'player_stat',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'person_id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'season',
                        'orig' => 'season',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'stat',
                        'orig' => 'stat',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/people/{id}/stats',
                  'parts' => [
                    'people',
                    '{person_id}',
                    'stats',
                  ],
                  'rename' => [
                    'param' => [
                      'id' => 'person_id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'person_id',
                      'season',
                      'stat',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.stats`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'person',
              ],
            ],
          ],
        ],
        'roster' => [
          'fields' => [
            [
              'name' => 'jerseyNumber',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'person',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'position',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'roster',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'team_id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'season',
                        'orig' => 'season',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/teams/{id}/roster',
                  'parts' => [
                    'teams',
                    '{team_id}',
                    'roster',
                  ],
                  'rename' => [
                    'param' => [
                      'id' => 'team_id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'season',
                      'team_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.roster`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'team',
              ],
            ],
          ],
        ],
        'schedule' => [
          'fields' => [
            [
              'name' => 'date',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'games',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'totalEvents',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'totalGames',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'totalItems',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'totalMatches',
              'type' => '`$INTEGER`',
            ],
          ],
          'name' => 'schedule',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'end_date',
                        'orig' => 'end_date',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'season',
                        'orig' => 'season',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'start_date',
                        'orig' => 'start_date',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'team_id',
                        'orig' => 'team_id',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/schedule',
                  'parts' => [
                    'schedule',
                  ],
                  'select' => [
                    'exist' => [
                      'end_date',
                      'season',
                      'start_date',
                      'team_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.dates`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'standing' => [
          'fields' => [
            [
              'name' => 'conference',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'division',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'teamRecords',
              'type' => '`$ARRAY`',
            ],
          ],
          'name' => 'standing',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'season',
                        'orig' => 'season',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/standings',
                  'parts' => [
                    'standings',
                  ],
                  'select' => [
                    'exist' => [
                      'season',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.records`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'team' => [
          'fields' => [
            [
              'name' => 'abbreviation',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'conference',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'copyright',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'division',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'firstYearOfPlay',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'franchise',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'id',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'link',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'locationName',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'teamName',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'teams',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'venue',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'team',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'expand',
                        'orig' => 'expand',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'season',
                        'orig' => 'season',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/teams',
                  'parts' => [
                    'teams',
                  ],
                  'select' => [
                    'exist' => [
                      'expand',
                      'season',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.teams`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$INTEGER`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'expand',
                        'orig' => 'expand',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/teams/{id}',
                  'parts' => [
                    'teams',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'expand',
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return NhlApiDocumentationFeatures::make_feature($name);
    }
}
