// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamaccountpasswordpolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		reflect.TypeOf((*IamAccountPasswordPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowUsersToChangePassword", GoGetter: "AllowUsersToChangePassword"},
			_jsii_.MemberProperty{JsiiProperty: "allowUsersToChangePasswordInput", GoGetter: "AllowUsersToChangePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "expirePasswords", GoGetter: "ExpirePasswords"},
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
			_jsii_.MemberProperty{JsiiProperty: "hardExpiry", GoGetter: "HardExpiry"},
			_jsii_.MemberProperty{JsiiProperty: "hardExpiryInput", GoGetter: "HardExpiryInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxPasswordAge", GoGetter: "MaxPasswordAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxPasswordAgeInput", GoGetter: "MaxPasswordAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLength", GoGetter: "MinimumPasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLengthInput", GoGetter: "MinimumPasswordLengthInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordReusePrevention", GoGetter: "PasswordReusePrevention"},
			_jsii_.MemberProperty{JsiiProperty: "passwordReusePreventionInput", GoGetter: "PasswordReusePreventionInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requireLowercaseCharacters", GoGetter: "RequireLowercaseCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "requireLowercaseCharactersInput", GoGetter: "RequireLowercaseCharactersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireNumbers", GoGetter: "RequireNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "requireNumbersInput", GoGetter: "RequireNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireSymbols", GoGetter: "RequireSymbols"},
			_jsii_.MemberProperty{JsiiProperty: "requireSymbolsInput", GoGetter: "RequireSymbolsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireUppercaseCharacters", GoGetter: "RequireUppercaseCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "requireUppercaseCharactersInput", GoGetter: "RequireUppercaseCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUsersToChangePassword", GoMethod: "ResetAllowUsersToChangePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetHardExpiry", GoMethod: "ResetHardExpiry"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPasswordAge", GoMethod: "ResetMaxPasswordAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumPasswordLength", GoMethod: "ResetMinimumPasswordLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordReusePrevention", GoMethod: "ResetPasswordReusePrevention"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireLowercaseCharacters", GoMethod: "ResetRequireLowercaseCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireNumbers", GoMethod: "ResetRequireNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireSymbols", GoMethod: "ResetRequireSymbols"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireUppercaseCharacters", GoMethod: "ResetRequireUppercaseCharacters"},
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
			j := jsiiProxy_IamAccountPasswordPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicyConfig",
		reflect.TypeOf((*IamAccountPasswordPolicyConfig)(nil)).Elem(),
	)
}
