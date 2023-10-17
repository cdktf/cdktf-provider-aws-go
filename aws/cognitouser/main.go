// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouser

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		reflect.TypeOf((*CognitoUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "attributesInput", GoGetter: "AttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientMetadata", GoGetter: "ClientMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "clientMetadataInput", GoGetter: "ClientMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "desiredDeliveryMediums", GoGetter: "DesiredDeliveryMediums"},
			_jsii_.MemberProperty{JsiiProperty: "desiredDeliveryMediumsInput", GoGetter: "DesiredDeliveryMediumsInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceAliasCreation", GoGetter: "ForceAliasCreation"},
			_jsii_.MemberProperty{JsiiProperty: "forceAliasCreationInput", GoGetter: "ForceAliasCreationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedDate", GoGetter: "LastModifiedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "messageAction", GoGetter: "MessageAction"},
			_jsii_.MemberProperty{JsiiProperty: "messageActionInput", GoGetter: "MessageActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "mfaSettingList", GoGetter: "MfaSettingList"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMfaSetting", GoGetter: "PreferredMfaSetting"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttributes", GoMethod: "ResetAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientMetadata", GoMethod: "ResetClientMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetDesiredDeliveryMediums", GoMethod: "ResetDesiredDeliveryMediums"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceAliasCreation", GoMethod: "ResetForceAliasCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageAction", GoMethod: "ResetMessageAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetTemporaryPassword", GoMethod: "ResetTemporaryPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidationData", GoMethod: "ResetValidationData"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "sub", GoGetter: "Sub"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "temporaryPassword", GoGetter: "TemporaryPassword"},
			_jsii_.MemberProperty{JsiiProperty: "temporaryPasswordInput", GoGetter: "TemporaryPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "validationData", GoGetter: "ValidationData"},
			_jsii_.MemberProperty{JsiiProperty: "validationDataInput", GoGetter: "ValidationDataInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.cognitoUser.CognitoUserConfig",
		reflect.TypeOf((*CognitoUserConfig)(nil)).Elem(),
	)
}
