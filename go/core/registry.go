package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewConferenceEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewDivisionEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewGameEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewPlayerEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewPlayerStatEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewRosterEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewScheduleEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewStandingEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

var NewTeamEntityFunc func(client *NhlApiDocumentationSDK, entopts map[string]any) NhlApiDocumentationEntity

