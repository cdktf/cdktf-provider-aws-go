// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolclient

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClient",
		reflect.TypeOf((*CognitoUserPoolClient)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTokenValidity", GoGetter: "AccessTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenValidityInput", GoGetter: "AccessTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlows", GoGetter: "AllowedOauthFlows"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsInput", GoGetter: "AllowedOauthFlowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsUserPoolClient", GoGetter: "AllowedOauthFlowsUserPoolClient"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsUserPoolClientInput", GoGetter: "AllowedOauthFlowsUserPoolClientInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthScopes", GoGetter: "AllowedOauthScopes"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthScopesInput", GoGetter: "AllowedOauthScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsConfiguration", GoGetter: "AnalyticsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsConfigurationInput", GoGetter: "AnalyticsConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "authSessionValidity", GoGetter: "AuthSessionValidity"},
			_jsii_.MemberProperty{JsiiProperty: "authSessionValidityInput", GoGetter: "AuthSessionValidityInput"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrls", GoGetter: "CallbackUrls"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrlsInput", GoGetter: "CallbackUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRedirectUri", GoGetter: "DefaultRedirectUri"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRedirectUriInput", GoGetter: "DefaultRedirectUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enablePropagateAdditionalUserContextData", GoGetter: "EnablePropagateAdditionalUserContextData"},
			_jsii_.MemberProperty{JsiiProperty: "enablePropagateAdditionalUserContextDataInput", GoGetter: "EnablePropagateAdditionalUserContextDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableTokenRevocation", GoGetter: "EnableTokenRevocation"},
			_jsii_.MemberProperty{JsiiProperty: "enableTokenRevocationInput", GoGetter: "EnableTokenRevocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "explicitAuthFlows", GoGetter: "ExplicitAuthFlows"},
			_jsii_.MemberProperty{JsiiProperty: "explicitAuthFlowsInput", GoGetter: "ExplicitAuthFlowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "generateSecret", GoGetter: "GenerateSecret"},
			_jsii_.MemberProperty{JsiiProperty: "generateSecretInput", GoGetter: "GenerateSecretInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idTokenValidity", GoGetter: "IdTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "idTokenValidityInput", GoGetter: "IdTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logoutUrls", GoGetter: "LogoutUrls"},
			_jsii_.MemberProperty{JsiiProperty: "logoutUrlsInput", GoGetter: "LogoutUrlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preventUserExistenceErrors", GoGetter: "PreventUserExistenceErrors"},
			_jsii_.MemberProperty{JsiiProperty: "preventUserExistenceErrorsInput", GoGetter: "PreventUserExistenceErrorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAnalyticsConfiguration", GoMethod: "PutAnalyticsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRefreshTokenRotation", GoMethod: "PutRefreshTokenRotation"},
			_jsii_.MemberMethod{JsiiMethod: "putTokenValidityUnits", GoMethod: "PutTokenValidityUnits"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readAttributes", GoGetter: "ReadAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "readAttributesInput", GoGetter: "ReadAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenRotation", GoGetter: "RefreshTokenRotation"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenRotationInput", GoGetter: "RefreshTokenRotationInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenValidity", GoGetter: "RefreshTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenValidityInput", GoGetter: "RefreshTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessTokenValidity", GoMethod: "ResetAccessTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthFlows", GoMethod: "ResetAllowedOauthFlows"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthFlowsUserPoolClient", GoMethod: "ResetAllowedOauthFlowsUserPoolClient"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthScopes", GoMethod: "ResetAllowedOauthScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnalyticsConfiguration", GoMethod: "ResetAnalyticsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthSessionValidity", GoMethod: "ResetAuthSessionValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetCallbackUrls", GoMethod: "ResetCallbackUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRedirectUri", GoMethod: "ResetDefaultRedirectUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePropagateAdditionalUserContextData", GoMethod: "ResetEnablePropagateAdditionalUserContextData"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableTokenRevocation", GoMethod: "ResetEnableTokenRevocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetExplicitAuthFlows", GoMethod: "ResetExplicitAuthFlows"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerateSecret", GoMethod: "ResetGenerateSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdTokenValidity", GoMethod: "ResetIdTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogoutUrls", GoMethod: "ResetLogoutUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventUserExistenceErrors", GoMethod: "ResetPreventUserExistenceErrors"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadAttributes", GoMethod: "ResetReadAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenRotation", GoMethod: "ResetRefreshTokenRotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenValidity", GoMethod: "ResetRefreshTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportedIdentityProviders", GoMethod: "ResetSupportedIdentityProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenValidityUnits", GoMethod: "ResetTokenValidityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteAttributes", GoMethod: "ResetWriteAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "supportedIdentityProviders", GoGetter: "SupportedIdentityProviders"},
			_jsii_.MemberProperty{JsiiProperty: "supportedIdentityProvidersInput", GoGetter: "SupportedIdentityProvidersInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tokenValidityUnits", GoGetter: "TokenValidityUnits"},
			_jsii_.MemberProperty{JsiiProperty: "tokenValidityUnitsInput", GoGetter: "TokenValidityUnitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "writeAttributes", GoGetter: "WriteAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "writeAttributesInput", GoGetter: "WriteAttributesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClient{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientAnalyticsConfiguration",
		reflect.TypeOf((*CognitoUserPoolClientAnalyticsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientAnalyticsConfigurationList",
		reflect.TypeOf((*CognitoUserPoolClientAnalyticsConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientAnalyticsConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolClientAnalyticsConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArnInput", GoGetter: "ApplicationArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "applicationIdInput", GoGetter: "ApplicationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "externalId", GoGetter: "ExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdInput", GoGetter: "ExternalIdInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationArn", GoMethod: "ResetApplicationArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationId", GoMethod: "ResetApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalId", GoMethod: "ResetExternalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserDataShared", GoMethod: "ResetUserDataShared"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userDataShared", GoGetter: "UserDataShared"},
			_jsii_.MemberProperty{JsiiProperty: "userDataSharedInput", GoGetter: "UserDataSharedInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientConfig",
		reflect.TypeOf((*CognitoUserPoolClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientRefreshTokenRotation",
		reflect.TypeOf((*CognitoUserPoolClientRefreshTokenRotation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientRefreshTokenRotationList",
		reflect.TypeOf((*CognitoUserPoolClientRefreshTokenRotationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientRefreshTokenRotationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientRefreshTokenRotationOutputReference",
		reflect.TypeOf((*CognitoUserPoolClientRefreshTokenRotationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "feature", GoGetter: "Feature"},
			_jsii_.MemberProperty{JsiiProperty: "featureInput", GoGetter: "FeatureInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetRetryGracePeriodSeconds", GoMethod: "ResetRetryGracePeriodSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retryGracePeriodSeconds", GoGetter: "RetryGracePeriodSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "retryGracePeriodSecondsInput", GoGetter: "RetryGracePeriodSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientRefreshTokenRotationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientTokenValidityUnits",
		reflect.TypeOf((*CognitoUserPoolClientTokenValidityUnits)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientTokenValidityUnitsList",
		reflect.TypeOf((*CognitoUserPoolClientTokenValidityUnitsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientTokenValidityUnitsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUserPoolClient.CognitoUserPoolClientTokenValidityUnitsOutputReference",
		reflect.TypeOf((*CognitoUserPoolClientTokenValidityUnitsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessToken", GoGetter: "AccessToken"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenInput", GoGetter: "AccessTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "idToken", GoGetter: "IdToken"},
			_jsii_.MemberProperty{JsiiProperty: "idTokenInput", GoGetter: "IdTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "refreshToken", GoGetter: "RefreshToken"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenInput", GoGetter: "RefreshTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessToken", GoMethod: "ResetAccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdToken", GoMethod: "ResetIdToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshToken", GoMethod: "ResetRefreshToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
