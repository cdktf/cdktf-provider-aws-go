// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayIntegration.ApiGatewayIntegration",
		reflect.TypeOf((*ApiGatewayIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cacheKeyParameters", GoGetter: "CacheKeyParameters"},
			_jsii_.MemberProperty{JsiiProperty: "cacheKeyParametersInput", GoGetter: "CacheKeyParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacheNamespace", GoGetter: "CacheNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "cacheNamespaceInput", GoGetter: "CacheNamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionId", GoGetter: "ConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "connectionIdInput", GoGetter: "ConnectionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTypeInput", GoGetter: "ConnectionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandling", GoGetter: "ContentHandling"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandlingInput", GoGetter: "ContentHandlingInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsInput", GoGetter: "CredentialsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpMethod", GoGetter: "HttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "httpMethodInput", GoGetter: "HttpMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "integrationHttpMethod", GoGetter: "IntegrationHttpMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationHttpMethodInput", GoGetter: "IntegrationHttpMethodInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughBehavior", GoGetter: "PassthroughBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughBehaviorInput", GoGetter: "PassthroughBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTlsConfig", GoMethod: "PutTlsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestParameters", GoGetter: "RequestParameters"},
			_jsii_.MemberProperty{JsiiProperty: "requestParametersInput", GoGetter: "RequestParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestTemplates", GoGetter: "RequestTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "requestTemplatesInput", GoGetter: "RequestTemplatesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheKeyParameters", GoMethod: "ResetCacheKeyParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheNamespace", GoMethod: "ResetCacheNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionId", GoMethod: "ResetConnectionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionType", GoMethod: "ResetConnectionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentHandling", GoMethod: "ResetContentHandling"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentials", GoMethod: "ResetCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationHttpMethod", GoMethod: "ResetIntegrationHttpMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassthroughBehavior", GoMethod: "ResetPassthroughBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestParameters", GoMethod: "ResetRequestParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTemplates", GoMethod: "ResetRequestTemplates"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutMilliseconds", GoMethod: "ResetTimeoutMilliseconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsConfig", GoMethod: "ResetTlsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetUri", GoMethod: "ResetUri"},
			_jsii_.MemberProperty{JsiiProperty: "resourceId", GoGetter: "ResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdInput", GoGetter: "ResourceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "restApiId", GoGetter: "RestApiId"},
			_jsii_.MemberProperty{JsiiProperty: "restApiIdInput", GoGetter: "RestApiIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMilliseconds", GoGetter: "TimeoutMilliseconds"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMillisecondsInput", GoGetter: "TimeoutMillisecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsConfig", GoGetter: "TlsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "tlsConfigInput", GoGetter: "TlsConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
			_jsii_.MemberProperty{JsiiProperty: "uriInput", GoGetter: "UriInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayIntegration.ApiGatewayIntegrationConfig",
		reflect.TypeOf((*ApiGatewayIntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apiGatewayIntegration.ApiGatewayIntegrationTlsConfig",
		reflect.TypeOf((*ApiGatewayIntegrationTlsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apiGatewayIntegration.ApiGatewayIntegrationTlsConfigOutputReference",
		reflect.TypeOf((*ApiGatewayIntegrationTlsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "insecureSkipVerification", GoGetter: "InsecureSkipVerification"},
			_jsii_.MemberProperty{JsiiProperty: "insecureSkipVerificationInput", GoGetter: "InsecureSkipVerificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureSkipVerification", GoMethod: "ResetInsecureSkipVerification"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayIntegrationTlsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
