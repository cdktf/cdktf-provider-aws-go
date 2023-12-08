// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snsplatformapplication

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		reflect.TypeOf((*SnsPlatformApplication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applePlatformBundleId", GoGetter: "ApplePlatformBundleId"},
			_jsii_.MemberProperty{JsiiProperty: "applePlatformBundleIdInput", GoGetter: "ApplePlatformBundleIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "applePlatformTeamId", GoGetter: "ApplePlatformTeamId"},
			_jsii_.MemberProperty{JsiiProperty: "applePlatformTeamIdInput", GoGetter: "ApplePlatformTeamIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "eventDeliveryFailureTopicArn", GoGetter: "EventDeliveryFailureTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "eventDeliveryFailureTopicArnInput", GoGetter: "EventDeliveryFailureTopicArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointCreatedTopicArn", GoGetter: "EventEndpointCreatedTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointCreatedTopicArnInput", GoGetter: "EventEndpointCreatedTopicArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointDeletedTopicArn", GoGetter: "EventEndpointDeletedTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointDeletedTopicArnInput", GoGetter: "EventEndpointDeletedTopicArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointUpdatedTopicArn", GoGetter: "EventEndpointUpdatedTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndpointUpdatedTopicArnInput", GoGetter: "EventEndpointUpdatedTopicArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "failureFeedbackRoleArn", GoGetter: "FailureFeedbackRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "failureFeedbackRoleArnInput", GoGetter: "FailureFeedbackRoleArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "platformCredential", GoGetter: "PlatformCredential"},
			_jsii_.MemberProperty{JsiiProperty: "platformCredentialInput", GoGetter: "PlatformCredentialInput"},
			_jsii_.MemberProperty{JsiiProperty: "platformInput", GoGetter: "PlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "platformPrincipal", GoGetter: "PlatformPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "platformPrincipalInput", GoGetter: "PlatformPrincipalInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplePlatformBundleId", GoMethod: "ResetApplePlatformBundleId"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplePlatformTeamId", GoMethod: "ResetApplePlatformTeamId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventDeliveryFailureTopicArn", GoMethod: "ResetEventDeliveryFailureTopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventEndpointCreatedTopicArn", GoMethod: "ResetEventEndpointCreatedTopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventEndpointDeletedTopicArn", GoMethod: "ResetEventEndpointDeletedTopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventEndpointUpdatedTopicArn", GoMethod: "ResetEventEndpointUpdatedTopicArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailureFeedbackRoleArn", GoMethod: "ResetFailureFeedbackRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlatformPrincipal", GoMethod: "ResetPlatformPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuccessFeedbackRoleArn", GoMethod: "ResetSuccessFeedbackRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuccessFeedbackSampleRate", GoMethod: "ResetSuccessFeedbackSampleRate"},
			_jsii_.MemberProperty{JsiiProperty: "successFeedbackRoleArn", GoGetter: "SuccessFeedbackRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "successFeedbackRoleArnInput", GoGetter: "SuccessFeedbackRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "successFeedbackSampleRate", GoGetter: "SuccessFeedbackSampleRate"},
			_jsii_.MemberProperty{JsiiProperty: "successFeedbackSampleRateInput", GoGetter: "SuccessFeedbackSampleRateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SnsPlatformApplication{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplicationConfig",
		reflect.TypeOf((*SnsPlatformApplicationConfig)(nil)).Elem(),
	)
}
