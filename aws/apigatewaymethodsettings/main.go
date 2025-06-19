// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaymethodsettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettings",
		reflect.TypeOf((*ApiGatewayMethodSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "methodPath", GoGetter: "MethodPath"},
			_jsii_.MemberProperty{JsiiProperty: "methodPathInput", GoGetter: "MethodPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putSettings", GoMethod: "PutSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberProperty{JsiiProperty: "restApiId", GoGetter: "RestApiId"},
			_jsii_.MemberProperty{JsiiProperty: "restApiIdInput", GoGetter: "RestApiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberProperty{JsiiProperty: "settingsInput", GoGetter: "SettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "stageNameInput", GoGetter: "StageNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayMethodSettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettingsConfig",
		reflect.TypeOf((*ApiGatewayMethodSettingsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettingsSettings",
		reflect.TypeOf((*ApiGatewayMethodSettingsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayMethodSettings.ApiGatewayMethodSettingsSettingsOutputReference",
		reflect.TypeOf((*ApiGatewayMethodSettingsSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cacheDataEncrypted", GoGetter: "CacheDataEncrypted"},
			_jsii_.MemberProperty{JsiiProperty: "cacheDataEncryptedInput", GoGetter: "CacheDataEncryptedInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacheTtlInSeconds", GoGetter: "CacheTtlInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "cacheTtlInSecondsInput", GoGetter: "CacheTtlInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachingEnabled", GoGetter: "CachingEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cachingEnabledInput", GoGetter: "CachingEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataTraceEnabled", GoGetter: "DataTraceEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "dataTraceEnabledInput", GoGetter: "DataTraceEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevel", GoGetter: "LoggingLevel"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevelInput", GoGetter: "LoggingLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabled", GoGetter: "MetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabledInput", GoGetter: "MetricsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireAuthorizationForCacheControl", GoGetter: "RequireAuthorizationForCacheControl"},
			_jsii_.MemberProperty{JsiiProperty: "requireAuthorizationForCacheControlInput", GoGetter: "RequireAuthorizationForCacheControlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheDataEncrypted", GoMethod: "ResetCacheDataEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheTtlInSeconds", GoMethod: "ResetCacheTtlInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachingEnabled", GoMethod: "ResetCachingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataTraceEnabled", GoMethod: "ResetDataTraceEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingLevel", GoMethod: "ResetLoggingLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsEnabled", GoMethod: "ResetMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireAuthorizationForCacheControl", GoMethod: "ResetRequireAuthorizationForCacheControl"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingBurstLimit", GoMethod: "ResetThrottlingBurstLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetThrottlingRateLimit", GoMethod: "ResetThrottlingRateLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnauthorizedCacheControlHeaderStrategy", GoMethod: "ResetUnauthorizedCacheControlHeaderStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimit", GoGetter: "ThrottlingBurstLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingBurstLimitInput", GoGetter: "ThrottlingBurstLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimit", GoGetter: "ThrottlingRateLimit"},
			_jsii_.MemberProperty{JsiiProperty: "throttlingRateLimitInput", GoGetter: "ThrottlingRateLimitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unauthorizedCacheControlHeaderStrategy", GoGetter: "UnauthorizedCacheControlHeaderStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "unauthorizedCacheControlHeaderStrategyInput", GoGetter: "UnauthorizedCacheControlHeaderStrategyInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayMethodSettingsSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
