<?php
declare(strict_types=1);

// NhlApiDocumentation SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

NhlApiDocumentationUtility::setRegistrar(function (NhlApiDocumentationUtility $u): void {
    $u->clean = [NhlApiDocumentationClean::class, 'call'];
    $u->done = [NhlApiDocumentationDone::class, 'call'];
    $u->make_error = [NhlApiDocumentationMakeError::class, 'call'];
    $u->feature_add = [NhlApiDocumentationFeatureAdd::class, 'call'];
    $u->feature_hook = [NhlApiDocumentationFeatureHook::class, 'call'];
    $u->feature_init = [NhlApiDocumentationFeatureInit::class, 'call'];
    $u->fetcher = [NhlApiDocumentationFetcher::class, 'call'];
    $u->make_fetch_def = [NhlApiDocumentationMakeFetchDef::class, 'call'];
    $u->make_context = [NhlApiDocumentationMakeContext::class, 'call'];
    $u->make_options = [NhlApiDocumentationMakeOptions::class, 'call'];
    $u->make_request = [NhlApiDocumentationMakeRequest::class, 'call'];
    $u->make_response = [NhlApiDocumentationMakeResponse::class, 'call'];
    $u->make_result = [NhlApiDocumentationMakeResult::class, 'call'];
    $u->make_point = [NhlApiDocumentationMakePoint::class, 'call'];
    $u->make_spec = [NhlApiDocumentationMakeSpec::class, 'call'];
    $u->make_url = [NhlApiDocumentationMakeUrl::class, 'call'];
    $u->param = [NhlApiDocumentationParam::class, 'call'];
    $u->prepare_auth = [NhlApiDocumentationPrepareAuth::class, 'call'];
    $u->prepare_body = [NhlApiDocumentationPrepareBody::class, 'call'];
    $u->prepare_headers = [NhlApiDocumentationPrepareHeaders::class, 'call'];
    $u->prepare_method = [NhlApiDocumentationPrepareMethod::class, 'call'];
    $u->prepare_params = [NhlApiDocumentationPrepareParams::class, 'call'];
    $u->prepare_path = [NhlApiDocumentationPreparePath::class, 'call'];
    $u->prepare_query = [NhlApiDocumentationPrepareQuery::class, 'call'];
    $u->result_basic = [NhlApiDocumentationResultBasic::class, 'call'];
    $u->result_body = [NhlApiDocumentationResultBody::class, 'call'];
    $u->result_headers = [NhlApiDocumentationResultHeaders::class, 'call'];
    $u->transform_request = [NhlApiDocumentationTransformRequest::class, 'call'];
    $u->transform_response = [NhlApiDocumentationTransformResponse::class, 'call'];
});
