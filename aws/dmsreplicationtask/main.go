// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationtask

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.dmsReplicationTask.DmsReplicationTask",
		reflect.TypeOf((*DmsReplicationTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdcStartPosition", GoGetter: "CdcStartPosition"},
			_jsii_.MemberProperty{JsiiProperty: "cdcStartPositionInput", GoGetter: "CdcStartPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdcStartTime", GoGetter: "CdcStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "cdcStartTimeInput", GoGetter: "CdcStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "migrationType", GoGetter: "MigrationType"},
			_jsii_.MemberProperty{JsiiProperty: "migrationTypeInput", GoGetter: "MigrationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInstanceArn", GoGetter: "ReplicationInstanceArn"},
			_jsii_.MemberProperty{JsiiProperty: "replicationInstanceArnInput", GoGetter: "ReplicationInstanceArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskArn", GoGetter: "ReplicationTaskArn"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskId", GoGetter: "ReplicationTaskId"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskIdInput", GoGetter: "ReplicationTaskIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskSettings", GoGetter: "ReplicationTaskSettings"},
			_jsii_.MemberProperty{JsiiProperty: "replicationTaskSettingsInput", GoGetter: "ReplicationTaskSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCdcStartPosition", GoMethod: "ResetCdcStartPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetCdcStartTime", GoMethod: "ResetCdcStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplicationTaskSettings", GoMethod: "ResetReplicationTaskSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceIdentifier", GoMethod: "ResetResourceIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartReplicationTask", GoMethod: "ResetStartReplicationTask"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdentifier", GoGetter: "ResourceIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdentifierInput", GoGetter: "ResourceIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceEndpointArn", GoGetter: "SourceEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceEndpointArnInput", GoGetter: "SourceEndpointArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "startReplicationTask", GoGetter: "StartReplicationTask"},
			_jsii_.MemberProperty{JsiiProperty: "startReplicationTaskInput", GoGetter: "StartReplicationTaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tableMappings", GoGetter: "TableMappings"},
			_jsii_.MemberProperty{JsiiProperty: "tableMappingsInput", GoGetter: "TableMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetEndpointArn", GoGetter: "TargetEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "targetEndpointArnInput", GoGetter: "TargetEndpointArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DmsReplicationTask{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.dmsReplicationTask.DmsReplicationTaskConfig",
		reflect.TypeOf((*DmsReplicationTaskConfig)(nil)).Elem(),
	)
}
