// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawssecretsmanagerrandompassword

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPassword",
		reflect.TypeOf((*DataAwsSecretsmanagerRandomPassword)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharacters", GoGetter: "ExcludeCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCharactersInput", GoGetter: "ExcludeCharactersInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeLowercase", GoGetter: "ExcludeLowercase"},
			_jsii_.MemberProperty{JsiiProperty: "excludeLowercaseInput", GoGetter: "ExcludeLowercaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeNumbers", GoGetter: "ExcludeNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "excludeNumbersInput", GoGetter: "ExcludeNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludePunctuation", GoGetter: "ExcludePunctuation"},
			_jsii_.MemberProperty{JsiiProperty: "excludePunctuationInput", GoGetter: "ExcludePunctuationInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeUppercase", GoGetter: "ExcludeUppercase"},
			_jsii_.MemberProperty{JsiiProperty: "excludeUppercaseInput", GoGetter: "ExcludeUppercaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeSpace", GoGetter: "IncludeSpace"},
			_jsii_.MemberProperty{JsiiProperty: "includeSpaceInput", GoGetter: "IncludeSpaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordLength", GoGetter: "PasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "passwordLengthInput", GoGetter: "PasswordLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "randomPassword", GoGetter: "RandomPassword"},
			_jsii_.MemberProperty{JsiiProperty: "randomPasswordInput", GoGetter: "RandomPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requireEachIncludedType", GoGetter: "RequireEachIncludedType"},
			_jsii_.MemberProperty{JsiiProperty: "requireEachIncludedTypeInput", GoGetter: "RequireEachIncludedTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeCharacters", GoMethod: "ResetExcludeCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeLowercase", GoMethod: "ResetExcludeLowercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeNumbers", GoMethod: "ResetExcludeNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludePunctuation", GoMethod: "ResetExcludePunctuation"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeUppercase", GoMethod: "ResetExcludeUppercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSpace", GoMethod: "ResetIncludeSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordLength", GoMethod: "ResetPasswordLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetRandomPassword", GoMethod: "ResetRandomPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireEachIncludedType", GoMethod: "ResetRequireEachIncludedType"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsSecretsmanagerRandomPassword{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.dataAwsSecretsmanagerRandomPassword.DataAwsSecretsmanagerRandomPasswordConfig",
		reflect.TypeOf((*DataAwsSecretsmanagerRandomPasswordConfig)(nil)).Elem(),
	)
}
