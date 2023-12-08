// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2integration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2Integration",
		reflect.TypeOf((*Apigatewayv2Integration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiId", GoGetter: "ApiId"},
			_jsii_.MemberProperty{JsiiProperty: "apiIdInput", GoGetter: "ApiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionId", GoGetter: "ConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "connectionIdInput", GoGetter: "ConnectionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connectionType", GoGetter: "ConnectionType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTypeInput", GoGetter: "ConnectionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandlingStrategy", GoGetter: "ContentHandlingStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "contentHandlingStrategyInput", GoGetter: "ContentHandlingStrategyInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsArn", GoGetter: "CredentialsArn"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsArnInput", GoGetter: "CredentialsArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "integrationMethod", GoGetter: "IntegrationMethod"},
			_jsii_.MemberProperty{JsiiProperty: "integrationMethodInput", GoGetter: "IntegrationMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "integrationResponseSelectionExpression", GoGetter: "IntegrationResponseSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "integrationSubtype", GoGetter: "IntegrationSubtype"},
			_jsii_.MemberProperty{JsiiProperty: "integrationSubtypeInput", GoGetter: "IntegrationSubtypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "integrationType", GoGetter: "IntegrationType"},
			_jsii_.MemberProperty{JsiiProperty: "integrationTypeInput", GoGetter: "IntegrationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUri", GoGetter: "IntegrationUri"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUriInput", GoGetter: "IntegrationUriInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughBehavior", GoGetter: "PassthroughBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughBehaviorInput", GoGetter: "PassthroughBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersion", GoGetter: "PayloadFormatVersion"},
			_jsii_.MemberProperty{JsiiProperty: "payloadFormatVersionInput", GoGetter: "PayloadFormatVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putResponseParameters", GoMethod: "PutResponseParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putTlsConfig", GoMethod: "PutTlsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestParameters", GoGetter: "RequestParameters"},
			_jsii_.MemberProperty{JsiiProperty: "requestParametersInput", GoGetter: "RequestParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestTemplates", GoGetter: "RequestTemplates"},
			_jsii_.MemberProperty{JsiiProperty: "requestTemplatesInput", GoGetter: "RequestTemplatesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionId", GoMethod: "ResetConnectionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionType", GoMethod: "ResetConnectionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentHandlingStrategy", GoMethod: "ResetContentHandlingStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentialsArn", GoMethod: "ResetCredentialsArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationMethod", GoMethod: "ResetIntegrationMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationSubtype", GoMethod: "ResetIntegrationSubtype"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrationUri", GoMethod: "ResetIntegrationUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassthroughBehavior", GoMethod: "ResetPassthroughBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayloadFormatVersion", GoMethod: "ResetPayloadFormatVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestParameters", GoMethod: "ResetRequestParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTemplates", GoMethod: "ResetRequestTemplates"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseParameters", GoMethod: "ResetResponseParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetTemplateSelectionExpression", GoMethod: "ResetTemplateSelectionExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutMilliseconds", GoMethod: "ResetTimeoutMilliseconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsConfig", GoMethod: "ResetTlsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "responseParameters", GoGetter: "ResponseParameters"},
			_jsii_.MemberProperty{JsiiProperty: "responseParametersInput", GoGetter: "ResponseParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "templateSelectionExpression", GoGetter: "TemplateSelectionExpression"},
			_jsii_.MemberProperty{JsiiProperty: "templateSelectionExpressionInput", GoGetter: "TemplateSelectionExpressionInput"},
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
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2Integration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationConfig",
		reflect.TypeOf((*Apigatewayv2IntegrationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationResponseParameters",
		reflect.TypeOf((*Apigatewayv2IntegrationResponseParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationResponseParametersList",
		reflect.TypeOf((*Apigatewayv2IntegrationResponseParametersList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_Apigatewayv2IntegrationResponseParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationResponseParametersOutputReference",
		reflect.TypeOf((*Apigatewayv2IntegrationResponseParametersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mappings", GoGetter: "Mappings"},
			_jsii_.MemberProperty{JsiiProperty: "mappingsInput", GoGetter: "MappingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statusCode", GoGetter: "StatusCode"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodeInput", GoGetter: "StatusCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2IntegrationResponseParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationTlsConfig",
		reflect.TypeOf((*Apigatewayv2IntegrationTlsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.apigatewayv2Integration.Apigatewayv2IntegrationTlsConfigOutputReference",
		reflect.TypeOf((*Apigatewayv2IntegrationTlsConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerNameToVerify", GoMethod: "ResetServerNameToVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serverNameToVerify", GoGetter: "ServerNameToVerify"},
			_jsii_.MemberProperty{JsiiProperty: "serverNameToVerifyInput", GoGetter: "ServerNameToVerifyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Apigatewayv2IntegrationTlsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
