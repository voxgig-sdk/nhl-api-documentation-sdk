package voxgignhlapidocumentationsdk

import (
	"github.com/voxgig-sdk/nhl-api-documentation-sdk/core"
	"github.com/voxgig-sdk/nhl-api-documentation-sdk/entity"
	"github.com/voxgig-sdk/nhl-api-documentation-sdk/feature"
	_ "github.com/voxgig-sdk/nhl-api-documentation-sdk/utility"
)

// Type aliases preserve external API.
type NhlApiDocumentationSDK = core.NhlApiDocumentationSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type NhlApiDocumentationEntity = core.NhlApiDocumentationEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type NhlApiDocumentationError = core.NhlApiDocumentationError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewConferenceEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewConferenceEntity(client, entopts)
	}
	core.NewDivisionEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewDivisionEntity(client, entopts)
	}
	core.NewGameEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewGameEntity(client, entopts)
	}
	core.NewPlayerEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewPlayerEntity(client, entopts)
	}
	core.NewPlayerStatEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewPlayerStatEntity(client, entopts)
	}
	core.NewRosterEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewRosterEntity(client, entopts)
	}
	core.NewScheduleEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewScheduleEntity(client, entopts)
	}
	core.NewStandingEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewStandingEntity(client, entopts)
	}
	core.NewTeamEntityFunc = func(client *core.NhlApiDocumentationSDK, entopts map[string]any) core.NhlApiDocumentationEntity {
		return entity.NewTeamEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewNhlApiDocumentationSDK = core.NewNhlApiDocumentationSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
