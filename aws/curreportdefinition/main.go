// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package curreportdefinition

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.curReportDefinition.CurReportDefinition",
		reflect.TypeOf((*CurReportDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalArtifacts", GoGetter: "AdditionalArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "additionalArtifactsInput", GoGetter: "AdditionalArtifactsInput"},
			_jsii_.MemberProperty{JsiiProperty: "additionalSchemaElements", GoGetter: "AdditionalSchemaElements"},
			_jsii_.MemberProperty{JsiiProperty: "additionalSchemaElementsInput", GoGetter: "AdditionalSchemaElementsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "formatInput", GoGetter: "FormatInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "refreshClosedReports", GoGetter: "RefreshClosedReports"},
			_jsii_.MemberProperty{JsiiProperty: "refreshClosedReportsInput", GoGetter: "RefreshClosedReportsInput"},
			_jsii_.MemberProperty{JsiiProperty: "reportName", GoGetter: "ReportName"},
			_jsii_.MemberProperty{JsiiProperty: "reportNameInput", GoGetter: "ReportNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "reportVersioning", GoGetter: "ReportVersioning"},
			_jsii_.MemberProperty{JsiiProperty: "reportVersioningInput", GoGetter: "ReportVersioningInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalArtifacts", GoMethod: "ResetAdditionalArtifacts"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshClosedReports", GoMethod: "ResetRefreshClosedReports"},
			_jsii_.MemberMethod{JsiiMethod: "resetReportVersioning", GoMethod: "ResetReportVersioning"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Prefix", GoMethod: "ResetS3Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInput", GoGetter: "S3BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Prefix", GoGetter: "S3Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "s3PrefixInput", GoGetter: "S3PrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3Region", GoGetter: "S3Region"},
			_jsii_.MemberProperty{JsiiProperty: "s3RegionInput", GoGetter: "S3RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnit", GoGetter: "TimeUnit"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnitInput", GoGetter: "TimeUnitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CurReportDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.curReportDefinition.CurReportDefinitionConfig",
		reflect.TypeOf((*CurReportDefinitionConfig)(nil)).Elem(),
	)
}
