// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsneptuneengineversion

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		reflect.TypeOf((*DataAwsNeptuneEngineVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCharacterSet", GoGetter: "DefaultCharacterSet"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOnly", GoGetter: "DefaultOnly"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOnlyInput", GoGetter: "DefaultOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "engineDescription", GoGetter: "EngineDescription"},
			_jsii_.MemberProperty{JsiiProperty: "engineInput", GoGetter: "EngineInput"},
			_jsii_.MemberProperty{JsiiProperty: "exportableLogTypes", GoGetter: "ExportableLogTypes"},
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
			_jsii_.MemberProperty{JsiiProperty: "hasMajorTarget", GoGetter: "HasMajorTarget"},
			_jsii_.MemberProperty{JsiiProperty: "hasMajorTargetInput", GoGetter: "HasMajorTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasMinorTarget", GoGetter: "HasMinorTarget"},
			_jsii_.MemberProperty{JsiiProperty: "hasMinorTargetInput", GoGetter: "HasMinorTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latest", GoGetter: "Latest"},
			_jsii_.MemberProperty{JsiiProperty: "latestInput", GoGetter: "LatestInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroupFamily", GoGetter: "ParameterGroupFamily"},
			_jsii_.MemberProperty{JsiiProperty: "parameterGroupFamilyInput", GoGetter: "ParameterGroupFamilyInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMajorTargets", GoGetter: "PreferredMajorTargets"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMajorTargetsInput", GoGetter: "PreferredMajorTargetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredUpgradeTargets", GoGetter: "PreferredUpgradeTargets"},
			_jsii_.MemberProperty{JsiiProperty: "preferredUpgradeTargetsInput", GoGetter: "PreferredUpgradeTargetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredVersions", GoGetter: "PreferredVersions"},
			_jsii_.MemberProperty{JsiiProperty: "preferredVersionsInput", GoGetter: "PreferredVersionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultOnly", GoMethod: "ResetDefaultOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetEngine", GoMethod: "ResetEngine"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasMajorTarget", GoMethod: "ResetHasMajorTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasMinorTarget", GoMethod: "ResetHasMinorTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLatest", GoMethod: "ResetLatest"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameterGroupFamily", GoMethod: "ResetParameterGroupFamily"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredMajorTargets", GoMethod: "ResetPreferredMajorTargets"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredUpgradeTargets", GoMethod: "ResetPreferredUpgradeTargets"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredVersions", GoMethod: "ResetPreferredVersions"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberProperty{JsiiProperty: "supportedCharacterSets", GoGetter: "SupportedCharacterSets"},
			_jsii_.MemberProperty{JsiiProperty: "supportedTimezones", GoGetter: "SupportedTimezones"},
			_jsii_.MemberProperty{JsiiProperty: "supportsGlobalDatabases", GoGetter: "SupportsGlobalDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "supportsLogExportsToCloudwatch", GoGetter: "SupportsLogExportsToCloudwatch"},
			_jsii_.MemberProperty{JsiiProperty: "supportsReadReplica", GoGetter: "SupportsReadReplica"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "validMajorTargets", GoGetter: "ValidMajorTargets"},
			_jsii_.MemberProperty{JsiiProperty: "validMinorTargets", GoGetter: "ValidMinorTargets"},
			_jsii_.MemberProperty{JsiiProperty: "validUpgradeTargets", GoGetter: "ValidUpgradeTargets"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionActual", GoGetter: "VersionActual"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsNeptuneEngineVersion{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersionConfig",
		reflect.TypeOf((*DataAwsNeptuneEngineVersionConfig)(nil)).Elem(),
	)
}
