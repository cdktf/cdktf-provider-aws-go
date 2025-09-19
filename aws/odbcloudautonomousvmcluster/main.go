// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudautonomousvmcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmCluster",
		reflect.TypeOf((*OdbCloudAutonomousVmCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousDataStoragePercentage", GoGetter: "AutonomousDataStoragePercentage"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousDataStorageSizeInTbs", GoGetter: "AutonomousDataStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "autonomousDataStorageSizeInTbsInput", GoGetter: "AutonomousDataStorageSizeInTbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "availableAutonomousDataStorageSizeInTbs", GoGetter: "AvailableAutonomousDataStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "availableContainerDatabases", GoGetter: "AvailableContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "availableCpus", GoGetter: "AvailableCpus"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureId", GoGetter: "CloudExadataInfrastructureId"},
			_jsii_.MemberProperty{JsiiProperty: "cloudExadataInfrastructureIdInput", GoGetter: "CloudExadataInfrastructureIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeModel", GoGetter: "ComputeModel"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCount", GoGetter: "CpuCoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountPerNode", GoGetter: "CpuCoreCountPerNode"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCoreCountPerNodeInput", GoGetter: "CpuCoreCountPerNodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cpuPercentage", GoGetter: "CpuPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInGbs", GoGetter: "DataStorageSizeInGbs"},
			_jsii_.MemberProperty{JsiiProperty: "dataStorageSizeInTbs", GoGetter: "DataStorageSizeInTbs"},
			_jsii_.MemberProperty{JsiiProperty: "dbServers", GoGetter: "DbServers"},
			_jsii_.MemberProperty{JsiiProperty: "dbServersInput", GoGetter: "DbServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "exadataStorageInTbsLowestScaledValue", GoGetter: "ExadataStorageInTbsLowestScaledValue"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isMtlsEnabledVmCluster", GoGetter: "IsMtlsEnabledVmCluster"},
			_jsii_.MemberProperty{JsiiProperty: "isMtlsEnabledVmClusterInput", GoGetter: "IsMtlsEnabledVmClusterInput"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModel", GoGetter: "LicenseModel"},
			_jsii_.MemberProperty{JsiiProperty: "licenseModelInput", GoGetter: "LicenseModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindow", GoGetter: "MaintenanceWindow"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowInput", GoGetter: "MaintenanceWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxAcdsLowestScaledValue", GoGetter: "MaxAcdsLowestScaledValue"},
			_jsii_.MemberProperty{JsiiProperty: "memoryPerOracleComputeUnitInGbs", GoGetter: "MemoryPerOracleComputeUnitInGbs"},
			_jsii_.MemberProperty{JsiiProperty: "memoryPerOracleComputeUnitInGbsInput", GoGetter: "MemoryPerOracleComputeUnitInGbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "memorySizeInGbs", GoGetter: "MemorySizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCount", GoGetter: "NodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "nonProvisionableAutonomousContainerDatabases", GoGetter: "NonProvisionableAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "ocid", GoGetter: "Ocid"},
			_jsii_.MemberProperty{JsiiProperty: "ociResourceAnchorName", GoGetter: "OciResourceAnchorName"},
			_jsii_.MemberProperty{JsiiProperty: "ociUrl", GoGetter: "OciUrl"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkId", GoGetter: "OdbNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "odbNetworkIdInput", GoGetter: "OdbNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "odbNodeStorageSizeInGbs", GoGetter: "OdbNodeStorageSizeInGbs"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "percentProgress", GoGetter: "PercentProgress"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisionableAutonomousContainerDatabases", GoGetter: "ProvisionableAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedAutonomousContainerDatabases", GoGetter: "ProvisionedAutonomousContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedCpus", GoGetter: "ProvisionedCpus"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceWindow", GoMethod: "PutMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reclaimableCpus", GoGetter: "ReclaimableCpus"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "reservedCpus", GoGetter: "ReservedCpus"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsMtlsEnabledVmCluster", GoMethod: "ResetIsMtlsEnabledVmCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseModel", GoMethod: "ResetLicenseModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindow", GoMethod: "ResetMaintenanceWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeZone", GoMethod: "ResetTimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortNonTls", GoGetter: "ScanListenerPortNonTls"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortNonTlsInput", GoGetter: "ScanListenerPortNonTlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTls", GoGetter: "ScanListenerPortTls"},
			_jsii_.MemberProperty{JsiiProperty: "scanListenerPortTlsInput", GoGetter: "ScanListenerPortTlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "shape", GoGetter: "Shape"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusReason", GoGetter: "StatusReason"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeDatabaseSslCertificateExpires", GoGetter: "TimeDatabaseSslCertificateExpires"},
			_jsii_.MemberProperty{JsiiProperty: "timeOrdsCertificateExpires", GoGetter: "TimeOrdsCertificateExpires"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "timeZoneInput", GoGetter: "TimeZoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalContainerDatabases", GoGetter: "TotalContainerDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "totalContainerDatabasesInput", GoGetter: "TotalContainerDatabasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudAutonomousVmCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterConfig",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindow",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeek",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeek)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList)(nil)).Elem(),
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
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekOutputReference",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowDaysOfWeekOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowList",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowList)(nil)).Elem(),
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
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowMonths",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowMonths)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList)(nil)).Elem(),
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
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowMonthsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowMonthsOutputReference",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowMonthsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowMonthsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "daysOfWeek", GoGetter: "DaysOfWeek"},
			_jsii_.MemberProperty{JsiiProperty: "daysOfWeekInput", GoGetter: "DaysOfWeekInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hoursOfDay", GoGetter: "HoursOfDay"},
			_jsii_.MemberProperty{JsiiProperty: "hoursOfDayInput", GoGetter: "HoursOfDayInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "leadTimeInWeeks", GoGetter: "LeadTimeInWeeks"},
			_jsii_.MemberProperty{JsiiProperty: "leadTimeInWeeksInput", GoGetter: "LeadTimeInWeeksInput"},
			_jsii_.MemberProperty{JsiiProperty: "months", GoGetter: "Months"},
			_jsii_.MemberProperty{JsiiProperty: "monthsInput", GoGetter: "MonthsInput"},
			_jsii_.MemberProperty{JsiiProperty: "preference", GoGetter: "Preference"},
			_jsii_.MemberProperty{JsiiProperty: "preferenceInput", GoGetter: "PreferenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDaysOfWeek", GoMethod: "PutDaysOfWeek"},
			_jsii_.MemberMethod{JsiiMethod: "putMonths", GoMethod: "PutMonths"},
			_jsii_.MemberMethod{JsiiMethod: "resetDaysOfWeek", GoMethod: "ResetDaysOfWeek"},
			_jsii_.MemberMethod{JsiiMethod: "resetHoursOfDay", GoMethod: "ResetHoursOfDay"},
			_jsii_.MemberMethod{JsiiMethod: "resetLeadTimeInWeeks", GoMethod: "ResetLeadTimeInWeeks"},
			_jsii_.MemberMethod{JsiiMethod: "resetMonths", GoMethod: "ResetMonths"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeeksOfMonth", GoMethod: "ResetWeeksOfMonth"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weeksOfMonth", GoGetter: "WeeksOfMonth"},
			_jsii_.MemberProperty{JsiiProperty: "weeksOfMonthInput", GoGetter: "WeeksOfMonthInput"},
		},
		func() interface{} {
			j := jsiiProxy_OdbCloudAutonomousVmClusterMaintenanceWindowOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterTimeouts",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.odbCloudAutonomousVmCluster.OdbCloudAutonomousVmClusterTimeoutsOutputReference",
		reflect.TypeOf((*OdbCloudAutonomousVmClusterTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_OdbCloudAutonomousVmClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
