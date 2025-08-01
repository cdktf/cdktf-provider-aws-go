// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		reflect.TypeOf((*FsxLustreFileSystem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoImportPolicy", GoGetter: "AutoImportPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoImportPolicyInput", GoGetter: "AutoImportPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "automaticBackupRetentionDays", GoGetter: "AutomaticBackupRetentionDays"},
			_jsii_.MemberProperty{JsiiProperty: "automaticBackupRetentionDaysInput", GoGetter: "AutomaticBackupRetentionDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupId", GoGetter: "BackupId"},
			_jsii_.MemberProperty{JsiiProperty: "backupIdInput", GoGetter: "BackupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToBackups", GoGetter: "CopyTagsToBackups"},
			_jsii_.MemberProperty{JsiiProperty: "copyTagsToBackupsInput", GoGetter: "CopyTagsToBackupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dailyAutomaticBackupStartTime", GoGetter: "DailyAutomaticBackupStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "dailyAutomaticBackupStartTimeInput", GoGetter: "DailyAutomaticBackupStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataCompressionType", GoGetter: "DataCompressionType"},
			_jsii_.MemberProperty{JsiiProperty: "dataCompressionTypeInput", GoGetter: "DataCompressionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataReadCacheConfiguration", GoGetter: "DataReadCacheConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "dataReadCacheConfigurationInput", GoGetter: "DataReadCacheConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentType", GoGetter: "DeploymentType"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentTypeInput", GoGetter: "DeploymentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsName", GoGetter: "DnsName"},
			_jsii_.MemberProperty{JsiiProperty: "driveCacheType", GoGetter: "DriveCacheType"},
			_jsii_.MemberProperty{JsiiProperty: "driveCacheTypeInput", GoGetter: "DriveCacheTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "efaEnabled", GoGetter: "EfaEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "efaEnabledInput", GoGetter: "EfaEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "exportPath", GoGetter: "ExportPath"},
			_jsii_.MemberProperty{JsiiProperty: "exportPathInput", GoGetter: "ExportPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemTypeVersion", GoGetter: "FileSystemTypeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemTypeVersionInput", GoGetter: "FileSystemTypeVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "finalBackupTags", GoGetter: "FinalBackupTags"},
			_jsii_.MemberProperty{JsiiProperty: "finalBackupTagsInput", GoGetter: "FinalBackupTagsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "importedFileChunkSize", GoGetter: "ImportedFileChunkSize"},
			_jsii_.MemberProperty{JsiiProperty: "importedFileChunkSizeInput", GoGetter: "ImportedFileChunkSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "importPath", GoGetter: "ImportPath"},
			_jsii_.MemberProperty{JsiiProperty: "importPathInput", GoGetter: "ImportPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logConfiguration", GoGetter: "LogConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logConfigurationInput", GoGetter: "LogConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "metadataConfiguration", GoGetter: "MetadataConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "metadataConfigurationInput", GoGetter: "MetadataConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "mountName", GoGetter: "MountName"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceIds", GoGetter: "NetworkInterfaceIds"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ownerId", GoGetter: "OwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "perUnitStorageThroughput", GoGetter: "PerUnitStorageThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "perUnitStorageThroughputInput", GoGetter: "PerUnitStorageThroughputInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDataReadCacheConfiguration", GoMethod: "PutDataReadCacheConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putLogConfiguration", GoMethod: "PutLogConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putMetadataConfiguration", GoMethod: "PutMetadataConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putRootSquashConfiguration", GoMethod: "PutRootSquashConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoImportPolicy", GoMethod: "ResetAutoImportPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomaticBackupRetentionDays", GoMethod: "ResetAutomaticBackupRetentionDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupId", GoMethod: "ResetBackupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopyTagsToBackups", GoMethod: "ResetCopyTagsToBackups"},
			_jsii_.MemberMethod{JsiiMethod: "resetDailyAutomaticBackupStartTime", GoMethod: "ResetDailyAutomaticBackupStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataCompressionType", GoMethod: "ResetDataCompressionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataReadCacheConfiguration", GoMethod: "ResetDataReadCacheConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentType", GoMethod: "ResetDeploymentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDriveCacheType", GoMethod: "ResetDriveCacheType"},
			_jsii_.MemberMethod{JsiiMethod: "resetEfaEnabled", GoMethod: "ResetEfaEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportPath", GoMethod: "ResetExportPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemTypeVersion", GoMethod: "ResetFileSystemTypeVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetFinalBackupTags", GoMethod: "ResetFinalBackupTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetImportedFileChunkSize", GoMethod: "ResetImportedFileChunkSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetImportPath", GoMethod: "ResetImportPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogConfiguration", GoMethod: "ResetLogConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataConfiguration", GoMethod: "ResetMetadataConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerUnitStorageThroughput", GoMethod: "ResetPerUnitStorageThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootSquashConfiguration", GoMethod: "ResetRootSquashConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipFinalBackup", GoMethod: "ResetSkipFinalBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageCapacity", GoMethod: "ResetStorageCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageType", GoMethod: "ResetStorageType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetThroughputCapacity", GoMethod: "ResetThroughputCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeeklyMaintenanceStartTime", GoMethod: "ResetWeeklyMaintenanceStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "rootSquashConfiguration", GoGetter: "RootSquashConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "rootSquashConfigurationInput", GoGetter: "RootSquashConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipFinalBackup", GoGetter: "SkipFinalBackup"},
			_jsii_.MemberProperty{JsiiProperty: "skipFinalBackupInput", GoGetter: "SkipFinalBackupInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageCapacity", GoGetter: "StorageCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "storageCapacityInput", GoGetter: "StorageCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "storageTypeInput", GoGetter: "StorageTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "throughputCapacity", GoGetter: "ThroughputCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "throughputCapacityInput", GoGetter: "ThroughputCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceStartTime", GoGetter: "WeeklyMaintenanceStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceStartTimeInput", GoGetter: "WeeklyMaintenanceStartTimeInput"},
		},
		func() interface{} {
			j := jsiiProxy_FsxLustreFileSystem{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemConfig",
		reflect.TypeOf((*FsxLustreFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemDataReadCacheConfiguration",
		reflect.TypeOf((*FsxLustreFileSystemDataReadCacheConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemDataReadCacheConfigurationOutputReference",
		reflect.TypeOf((*FsxLustreFileSystemDataReadCacheConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetSize", GoMethod: "ResetSize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInput", GoGetter: "SizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sizingMode", GoGetter: "SizingMode"},
			_jsii_.MemberProperty{JsiiProperty: "sizingModeInput", GoGetter: "SizingModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FsxLustreFileSystemDataReadCacheConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemLogConfiguration",
		reflect.TypeOf((*FsxLustreFileSystemLogConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemLogConfigurationOutputReference",
		reflect.TypeOf((*FsxLustreFileSystemLogConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "destinationInput", GoGetter: "DestinationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "levelInput", GoGetter: "LevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestination", GoMethod: "ResetDestination"},
			_jsii_.MemberMethod{JsiiMethod: "resetLevel", GoMethod: "ResetLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FsxLustreFileSystemLogConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemMetadataConfiguration",
		reflect.TypeOf((*FsxLustreFileSystemMetadataConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemMetadataConfigurationOutputReference",
		reflect.TypeOf((*FsxLustreFileSystemMetadataConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "iopsInput", GoGetter: "IopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIops", GoMethod: "ResetIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetMode", GoMethod: "ResetMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FsxLustreFileSystemMetadataConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemRootSquashConfiguration",
		reflect.TypeOf((*FsxLustreFileSystemRootSquashConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemRootSquashConfigurationOutputReference",
		reflect.TypeOf((*FsxLustreFileSystemRootSquashConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "noSquashNids", GoGetter: "NoSquashNids"},
			_jsii_.MemberProperty{JsiiProperty: "noSquashNidsInput", GoGetter: "NoSquashNidsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoSquashNids", GoMethod: "ResetNoSquashNids"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootSquash", GoMethod: "ResetRootSquash"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rootSquash", GoGetter: "RootSquash"},
			_jsii_.MemberProperty{JsiiProperty: "rootSquashInput", GoGetter: "RootSquashInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FsxLustreFileSystemRootSquashConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemTimeouts",
		reflect.TypeOf((*FsxLustreFileSystemTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystemTimeoutsOutputReference",
		reflect.TypeOf((*FsxLustreFileSystemTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
