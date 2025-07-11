// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsclusterinstance

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.rdsClusterInstance.RdsClusterInstance",
		reflect.TypeOf((*RdsClusterInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applyImmediately", GoGetter: "ApplyImmediately"},
			_jsii_.MemberProperty{JsiiProperty: "applyImmediatelyInput", GoGetter: "ApplyImmediatelyInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgrade", GoGetter: "AutoMinorVersionUpgrade"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgradeInput", GoGetter: "AutoMinorVersionUpgradeInput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneInput", GoGetter: "AvailabilityZoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertIdentifier", GoGetter: "CaCertIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "caCertIdentifierInput", GoGetter: "CaCertIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifierInput", GoGetter: "ClusterIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToSnapshot", GoGetter: "CopyTagsToSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToSnapshotInput", GoGetter: "CopyTagsToSnapshotInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customIamInstanceProfile", GoGetter: "CustomIamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "customIamInstanceProfileInput", GoGetter: "CustomIamInstanceProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbiResourceId", GoGetter: "DbiResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dbParameterGroupName", GoGetter: "DbParameterGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbParameterGroupNameInput", GoGetter: "DbParameterGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupName", GoGetter: "DbSubnetGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "dbSubnetGroupNameInput", GoGetter: "DbSubnetGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "engine", GoGetter: "Engine"},
			_jsii_.MemberProperty{JsiiProperty: "engineInput", GoGetter: "EngineInput"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersionActual", GoGetter: "EngineVersionActual"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersionInput", GoGetter: "EngineVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceDestroy", GoGetter: "ForceDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "forceDestroyInput", GoGetter: "ForceDestroyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identifier", GoGetter: "Identifier"},
			_jsii_.MemberProperty{JsiiProperty: "identifierInput", GoGetter: "IdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "identifierPrefix", GoGetter: "IdentifierPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "identifierPrefixInput", GoGetter: "IdentifierPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "instanceClass", GoGetter: "InstanceClass"},
			_jsii_.MemberProperty{JsiiProperty: "instanceClassInput", GoGetter: "InstanceClassInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringInterval", GoGetter: "MonitoringInterval"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringIntervalInput", GoGetter: "MonitoringIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringRoleArn", GoGetter: "MonitoringRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringRoleArnInput", GoGetter: "MonitoringRoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "networkType", GoGetter: "NetworkType"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsEnabled", GoGetter: "PerformanceInsightsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsEnabledInput", GoGetter: "PerformanceInsightsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsKmsKeyId", GoGetter: "PerformanceInsightsKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsKmsKeyIdInput", GoGetter: "PerformanceInsightsKmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsRetentionPeriod", GoGetter: "PerformanceInsightsRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "performanceInsightsRetentionPeriodInput", GoGetter: "PerformanceInsightsRetentionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "preferredBackupWindow", GoGetter: "PreferredBackupWindow"},
			_jsii_.MemberProperty{JsiiProperty: "preferredBackupWindowInput", GoGetter: "PreferredBackupWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMaintenanceWindow", GoGetter: "PreferredMaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "preferredMaintenanceWindowInput", GoGetter: "PreferredMaintenanceWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "promotionTier", GoGetter: "PromotionTier"},
			_jsii_.MemberProperty{JsiiProperty: "promotionTierInput", GoGetter: "PromotionTierInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessibleInput", GoGetter: "PubliclyAccessibleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplyImmediately", GoMethod: "ResetApplyImmediately"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoMinorVersionUpgrade", GoMethod: "ResetAutoMinorVersionUpgrade"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailabilityZone", GoMethod: "ResetAvailabilityZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertIdentifier", GoMethod: "ResetCaCertIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyTagsToSnapshot", GoMethod: "ResetCopyTagsToSnapshot"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomIamInstanceProfile", GoMethod: "ResetCustomIamInstanceProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbParameterGroupName", GoMethod: "ResetDbParameterGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbSubnetGroupName", GoMethod: "ResetDbSubnetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEngineVersion", GoMethod: "ResetEngineVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceDestroy", GoMethod: "ResetForceDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentifier", GoMethod: "ResetIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentifierPrefix", GoMethod: "ResetIdentifierPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetMonitoringInterval", GoMethod: "ResetMonitoringInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetMonitoringRoleArn", GoMethod: "ResetMonitoringRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerformanceInsightsEnabled", GoMethod: "ResetPerformanceInsightsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerformanceInsightsKmsKeyId", GoMethod: "ResetPerformanceInsightsKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerformanceInsightsRetentionPeriod", GoMethod: "ResetPerformanceInsightsRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredBackupWindow", GoMethod: "ResetPreferredBackupWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredMaintenanceWindow", GoMethod: "ResetPreferredMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetPromotionTier", GoMethod: "ResetPromotionTier"},
			_jsii_.MemberMethod{JsiiMethod: "resetPubliclyAccessible", GoMethod: "ResetPubliclyAccessible"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "storageEncrypted", GoGetter: "StorageEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "writer", GoGetter: "Writer"},
		},
		func() interface{} {
			j := jsiiProxy_RdsClusterInstance{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.rdsClusterInstance.RdsClusterInstanceConfig",
		reflect.TypeOf((*RdsClusterInstanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.rdsClusterInstance.RdsClusterInstanceTimeouts",
		reflect.TypeOf((*RdsClusterInstanceTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.rdsClusterInstance.RdsClusterInstanceTimeoutsOutputReference",
		reflect.TypeOf((*RdsClusterInstanceTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_RdsClusterInstanceTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
