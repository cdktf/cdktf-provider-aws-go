// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/fsxlustrefilesystem/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system}.
type FsxLustreFileSystem interface {
	cdktf.TerraformResource
	Arn() *string
	AutoImportPolicy() *string
	SetAutoImportPolicy(val *string)
	AutoImportPolicyInput() *string
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToBackups() interface{}
	SetCopyTagsToBackups(val interface{})
	CopyTagsToBackupsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DailyAutomaticBackupStartTime() *string
	SetDailyAutomaticBackupStartTime(val *string)
	DailyAutomaticBackupStartTimeInput() *string
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	DataReadCacheConfiguration() FsxLustreFileSystemDataReadCacheConfigurationOutputReference
	DataReadCacheConfigurationInput() *FsxLustreFileSystemDataReadCacheConfiguration
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DnsName() *string
	DriveCacheType() *string
	SetDriveCacheType(val *string)
	DriveCacheTypeInput() *string
	EfaEnabled() interface{}
	SetEfaEnabled(val interface{})
	EfaEnabledInput() interface{}
	ExportPath() *string
	SetExportPath(val *string)
	ExportPathInput() *string
	FileSystemTypeVersion() *string
	SetFileSystemTypeVersion(val *string)
	FileSystemTypeVersionInput() *string
	FinalBackupTags() *map[string]*string
	SetFinalBackupTags(val *map[string]*string)
	FinalBackupTagsInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImportedFileChunkSize() *float64
	SetImportedFileChunkSize(val *float64)
	ImportedFileChunkSizeInput() *float64
	ImportPath() *string
	SetImportPath(val *string)
	ImportPathInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfiguration() FsxLustreFileSystemLogConfigurationOutputReference
	LogConfigurationInput() *FsxLustreFileSystemLogConfiguration
	MetadataConfiguration() FsxLustreFileSystemMetadataConfigurationOutputReference
	MetadataConfigurationInput() *FsxLustreFileSystemMetadataConfiguration
	MountName() *string
	NetworkInterfaceIds() *[]*string
	// The tree node.
	Node() constructs.Node
	OwnerId() *string
	PerUnitStorageThroughput() *float64
	SetPerUnitStorageThroughput(val *float64)
	PerUnitStorageThroughputInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RootSquashConfiguration() FsxLustreFileSystemRootSquashConfigurationOutputReference
	RootSquashConfigurationInput() *FsxLustreFileSystemRootSquashConfiguration
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SkipFinalBackup() interface{}
	SetSkipFinalBackup(val interface{})
	SkipFinalBackupInput() interface{}
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageCapacityInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThroughputCapacity() *float64
	SetThroughputCapacity(val *float64)
	ThroughputCapacityInput() *float64
	Timeouts() FsxLustreFileSystemTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	WeeklyMaintenanceStartTime() *string
	SetWeeklyMaintenanceStartTime(val *string)
	WeeklyMaintenanceStartTimeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDataReadCacheConfiguration(value *FsxLustreFileSystemDataReadCacheConfiguration)
	PutLogConfiguration(value *FsxLustreFileSystemLogConfiguration)
	PutMetadataConfiguration(value *FsxLustreFileSystemMetadataConfiguration)
	PutRootSquashConfiguration(value *FsxLustreFileSystemRootSquashConfiguration)
	PutTimeouts(value *FsxLustreFileSystemTimeouts)
	ResetAutoImportPolicy()
	ResetAutomaticBackupRetentionDays()
	ResetBackupId()
	ResetCopyTagsToBackups()
	ResetDailyAutomaticBackupStartTime()
	ResetDataCompressionType()
	ResetDataReadCacheConfiguration()
	ResetDeploymentType()
	ResetDriveCacheType()
	ResetEfaEnabled()
	ResetExportPath()
	ResetFileSystemTypeVersion()
	ResetFinalBackupTags()
	ResetId()
	ResetImportedFileChunkSize()
	ResetImportPath()
	ResetKmsKeyId()
	ResetLogConfiguration()
	ResetMetadataConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerUnitStorageThroughput()
	ResetRegion()
	ResetRootSquashConfiguration()
	ResetSecurityGroupIds()
	ResetSkipFinalBackup()
	ResetStorageCapacity()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetThroughputCapacity()
	ResetTimeouts()
	ResetWeeklyMaintenanceStartTime()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxLustreFileSystem
type jsiiProxy_FsxLustreFileSystem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxLustreFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutoImportPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutoImportPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataReadCacheConfiguration() FsxLustreFileSystemDataReadCacheConfigurationOutputReference {
	var returns FsxLustreFileSystemDataReadCacheConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataReadCacheConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataReadCacheConfigurationInput() *FsxLustreFileSystemDataReadCacheConfiguration {
	var returns *FsxLustreFileSystemDataReadCacheConfiguration
	_jsii_.Get(
		j,
		"dataReadCacheConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DriveCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DriveCacheTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) EfaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) EfaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FileSystemTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) LogConfiguration() FsxLustreFileSystemLogConfigurationOutputReference {
	var returns FsxLustreFileSystemLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) LogConfigurationInput() *FsxLustreFileSystemLogConfiguration {
	var returns *FsxLustreFileSystemLogConfiguration
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) MetadataConfiguration() FsxLustreFileSystemMetadataConfigurationOutputReference {
	var returns FsxLustreFileSystemMetadataConfigurationOutputReference
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) MetadataConfigurationInput() *FsxLustreFileSystemMetadataConfiguration {
	var returns *FsxLustreFileSystemMetadataConfiguration
	_jsii_.Get(
		j,
		"metadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) RootSquashConfiguration() FsxLustreFileSystemRootSquashConfigurationOutputReference {
	var returns FsxLustreFileSystemRootSquashConfigurationOutputReference
	_jsii_.Get(
		j,
		"rootSquashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) RootSquashConfigurationInput() *FsxLustreFileSystemRootSquashConfiguration {
	var returns *FsxLustreFileSystemRootSquashConfiguration
	_jsii_.Get(
		j,
		"rootSquashConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Timeouts() FsxLustreFileSystemTimeoutsOutputReference {
	var returns FsxLustreFileSystemTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem(scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) FsxLustreFileSystem {
	_init_.Initialize()

	if err := validateNewFsxLustreFileSystemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxLustreFileSystem{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem_Override(f FsxLustreFileSystem, scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetAutoImportPolicy(val *string) {
	if err := j.validateSetAutoImportPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImportPolicy",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetAutomaticBackupRetentionDays(val *float64) {
	if err := j.validateSetAutomaticBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetDailyAutomaticBackupStartTime(val *string) {
	if err := j.validateSetDailyAutomaticBackupStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetDriveCacheType(val *string) {
	if err := j.validateSetDriveCacheTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driveCacheType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetEfaEnabled(val interface{}) {
	if err := j.validateSetEfaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efaEnabled",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetExportPath(val *string) {
	if err := j.validateSetExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetFileSystemTypeVersion(val *string) {
	if err := j.validateSetFileSystemTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetImportedFileChunkSize(val *float64) {
	if err := j.validateSetImportedFileChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetImportPath(val *string) {
	if err := j.validateSetImportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPath",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetThroughputCapacity(val *float64) {
	if err := j.validateSetThroughputCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Generates CDKTF code for importing a FsxLustreFileSystem resource upon running "cdktf plan <stack-name>".
func FsxLustreFileSystem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFsxLustreFileSystem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FsxLustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxLustreFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxLustreFileSystem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxLustreFileSystem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxLustreFileSystem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxLustreFileSystem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxLustreFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.fsxLustreFileSystem.FsxLustreFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutDataReadCacheConfiguration(value *FsxLustreFileSystemDataReadCacheConfiguration) {
	if err := f.validatePutDataReadCacheConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDataReadCacheConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutLogConfiguration(value *FsxLustreFileSystemLogConfiguration) {
	if err := f.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutMetadataConfiguration(value *FsxLustreFileSystemMetadataConfiguration) {
	if err := f.validatePutMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putMetadataConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutRootSquashConfiguration(value *FsxLustreFileSystemRootSquashConfiguration) {
	if err := f.validatePutRootSquashConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putRootSquashConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutTimeouts(value *FsxLustreFileSystemTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		f,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		f,
		"resetBackupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDataReadCacheConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetDataReadCacheConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		f,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDriveCacheType() {
	_jsii_.InvokeVoid(
		f,
		"resetDriveCacheType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetEfaEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetEfaEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetExportPath() {
	_jsii_.InvokeVoid(
		f,
		"resetExportPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetFileSystemTypeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetFileSystemTypeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		f,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		f,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetImportPath() {
	_jsii_.InvokeVoid(
		f,
		"resetImportPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetMetadataConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetMetadataConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		f,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetRegion() {
	_jsii_.InvokeVoid(
		f,
		"resetRegion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetRootSquashConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetRootSquashConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetThroughputCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetThroughputCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

