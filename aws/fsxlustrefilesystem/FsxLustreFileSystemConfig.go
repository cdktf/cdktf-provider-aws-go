// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxLustreFileSystemConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#subnet_ids FsxLustreFileSystem#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#auto_import_policy FsxLustreFileSystem#auto_import_policy}.
	AutoImportPolicy *string `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#automatic_backup_retention_days FsxLustreFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#backup_id FsxLustreFileSystem#backup_id}.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#copy_tags_to_backups FsxLustreFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#daily_automatic_backup_start_time FsxLustreFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#data_compression_type FsxLustreFileSystem#data_compression_type}.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#deployment_type FsxLustreFileSystem#deployment_type}.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#drive_cache_type FsxLustreFileSystem#drive_cache_type}.
	DriveCacheType *string `field:"optional" json:"driveCacheType" yaml:"driveCacheType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#export_path FsxLustreFileSystem#export_path}.
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#file_system_type_version FsxLustreFileSystem#file_system_type_version}.
	FileSystemTypeVersion *string `field:"optional" json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#final_backup_tags FsxLustreFileSystem#final_backup_tags}.
	FinalBackupTags *map[string]*string `field:"optional" json:"finalBackupTags" yaml:"finalBackupTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#id FsxLustreFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#imported_file_chunk_size FsxLustreFileSystem#imported_file_chunk_size}.
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#import_path FsxLustreFileSystem#import_path}.
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#kms_key_id FsxLustreFileSystem#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#log_configuration FsxLustreFileSystem#log_configuration}
	LogConfiguration *FsxLustreFileSystemLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#metadata_configuration FsxLustreFileSystem#metadata_configuration}
	MetadataConfiguration *FsxLustreFileSystemMetadataConfiguration `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#per_unit_storage_throughput FsxLustreFileSystem#per_unit_storage_throughput}.
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// root_squash_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#root_squash_configuration FsxLustreFileSystem#root_squash_configuration}
	RootSquashConfiguration *FsxLustreFileSystemRootSquashConfiguration `field:"optional" json:"rootSquashConfiguration" yaml:"rootSquashConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#security_group_ids FsxLustreFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#skip_final_backup FsxLustreFileSystem#skip_final_backup}.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#storage_capacity FsxLustreFileSystem#storage_capacity}.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#storage_type FsxLustreFileSystem#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#tags FsxLustreFileSystem#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#tags_all FsxLustreFileSystem#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#timeouts FsxLustreFileSystem#timeouts}
	Timeouts *FsxLustreFileSystemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/fsx_lustre_file_system#weekly_maintenance_start_time FsxLustreFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

