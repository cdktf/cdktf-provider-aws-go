// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmCluster",
		reflect.TypeOf((*OdbCloudVmCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureId", GoGetter: "CloudExadataInfrastructureId"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureIdInput", GoGetter: "CloudExadataInfrastructureIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeModel", GoGetter: "ComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountInput", GoGetter: "CpuCoreCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dataCollectionOptions", GoGetter: "DataCollectionOptions"},
			_jsii_.MemberProperty{JsiiProperty: "dataCollectionOptionsInput", GoGetter: "DataCollectionOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTbs", GoGetter: "DataStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTbsInput", GoGetter: "DataStorageSizeInTbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbNodeStorageSizeInGbs", GoGetter: "DbNodeStorageSizeInGbs"},
			_jsii_.MemberProperty{JsiiProperty: "dbNodeStorageSizeInGbsInput", GoGetter: "DbNodeStorageSizeInGbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dbServers", GoGetter: "DbServers"},
			_jsii_.MemberProperty{JsiiProperty: "dbServersInput", GoGetter: "DbServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "diskRedundancy", GoGetter: "DiskRedundancy"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
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
			_jsii_.MemberProperty{JsiiProperty: "giVersion", GoGetter: "GiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "giVersionComputed", GoGetter: "GiVersionComputed"},
			_jsii_.MemberProperty{JsiiProperty: "giVersionInput", GoGetter: "GiVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hostnamePrefix", GoGetter: "HostnamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "hostnamePrefixComputed", GoGetter: "HostnamePrefixComputed"},
			_jsii_.MemberProperty{JsiiProperty: "hostnamePrefixInput", GoGetter: "HostnamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iormConfigCache", GoGetter: "IormConfigCache"},
			_jsii_.MemberProperty{JsiiProperty: "isLocalBackupEnabled", GoGetter: "IsLocalBackupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isLocalBackupEnabledInput", GoGetter: "IsLocalBackupEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSparseDiskgroupEnabled", GoGetter: "IsSparseDiskgroupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isSparseDiskgroupEnabledInput", GoGetter: "IsSparseDiskgroupEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdateHistoryEntryId", GoGetter: "LastUpdateHistoryEntryId"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModelInput", GoGetter: "LicenseModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "listenerPort", GoGetter: "ListenerPort"},
			_jsii_.MemberProperty{JsiiProperty: "memorySizeInGbs", GoGetter: "MemorySizeInGbs"},
			_jsii_.MemberProperty{JsiiProperty: "memorySizeInGbsInput", GoGetter: "MemorySizeInGbsInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCount", GoGetter: "NodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "ocid", GoGetter: "Ocid"},
			_jsii_.MemberProperty{JsiiProperty: "ociResourceAnchorName", GoGetter: "OciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "ociUrl", GoGetter: "OciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkId", GoGetter: "OdbNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkIdInput", GoGetter: "OdbNetworkIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "percentProgress", GoGetter: "PercentProgress"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDataCollectionOptions", GoMethod: "PutDataCollectionOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterName", GoMethod: "ResetClusterName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataCollectionOptions", GoMethod: "ResetDataCollectionOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbNodeStorageSizeInGbs", GoMethod: "ResetDbNodeStorageSizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsLocalBackupEnabled", GoMethod: "ResetIsLocalBackupEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsSparseDiskgroupEnabled", GoMethod: "ResetIsSparseDiskgroupEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseModel", GoMethod: "ResetLicenseModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemorySizeInGbs", GoMethod: "ResetMemorySizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetScanListenerPortTcp", GoMethod: "ResetScanListenerPortTcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimezone", GoMethod: "ResetTimezone"},
			_jsii_.MemberProperty{JsiiProperty: "scanDnsName", GoGetter: "ScanDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "scanDnsRecordId", GoGetter: "ScanDnsRecordId"},
			_jsii_.MemberProperty{JsiiProperty: "scanIpIds", GoGetter: "ScanIpIds"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTcp", GoGetter: "ScanListenerPortTcp"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTcpInput", GoGetter: "ScanListenerPortTcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "shape", GoGetter: "Shape"},
			_jsii_.MemberProperty{JsiiProperty: "sshPublicKeys", GoGetter: "SshPublicKeys"},
			_jsii_.MemberProperty{JsiiProperty: "sshPublicKeysInput", GoGetter: "SshPublicKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberProperty{JsiiProperty: "storageSizeInGbs", GoGetter: "StorageSizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "systemVersion", GoGetter: "SystemVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timezone", GoGetter: "Timezone"},
			_jsii_.MemberProperty{JsiiProperty: "timezoneInput", GoGetter: "TimezoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vipIds", GoGetter: "VipIds"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterConfig",
		reflect.TypeOf((*OdbCloudVmClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterDataCollectionOptions",
		reflect.TypeOf((*OdbCloudVmClusterDataCollectionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterDataCollectionOptionsList",
		reflect.TypeOf((*OdbCloudVmClusterDataCollectionOptionsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_OdbCloudVmClusterDataCollectionOptionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterDataCollectionOptionsOutputReference",
		reflect.TypeOf((*OdbCloudVmClusterDataCollectionOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "isDiagnosticsEventsEnabled", GoGetter: "IsDiagnosticsEventsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isDiagnosticsEventsEnabledInput", GoGetter: "IsDiagnosticsEventsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isHealthMonitoringEnabled", GoGetter: "IsHealthMonitoringEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isHealthMonitoringEnabledInput", GoGetter: "IsHealthMonitoringEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "isIncidentLogsEnabled", GoGetter: "IsIncidentLogsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isIncidentLogsEnabledInput", GoGetter: "IsIncidentLogsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmClusterDataCollectionOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCache",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCacheDbPlans",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCacheDbPlans)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCacheDbPlansList",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCacheDbPlansList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmClusterIormConfigCacheDbPlansList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCacheDbPlansOutputReference",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCacheDbPlansOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbName", GoGetter: "DbName"},
			_jsii_.MemberProperty{JsiiProperty: "flashCacheLimit", GoGetter: "FlashCacheLimit"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "share", GoGetter: "Share"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmClusterIormConfigCacheDbPlansOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCacheList",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCacheList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmClusterIormConfigCacheList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterIormConfigCacheOutputReference",
		reflect.TypeOf((*OdbCloudVmClusterIormConfigCacheOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbPlans", GoGetter: "DbPlans"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycleDetails", GoGetter: "LifecycleDetails"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleState", GoGetter: "LifecycleState"},
			_jsii_.MemberProperty{JsiiProperty: "objective", GoGetter: "Objective"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudVmClusterIormConfigCacheOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterTimeouts",
		reflect.TypeOf((*OdbCloudVmClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudVmCluster.OdbCloudVmClusterTimeoutsOutputReference",
		reflect.TypeOf((*OdbCloudVmClusterTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_OdbCloudVmClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
