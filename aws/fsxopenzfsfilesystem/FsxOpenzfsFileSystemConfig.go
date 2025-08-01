// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsfilesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOpenzfsFileSystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#deployment_type FsxOpenzfsFileSystem#deployment_type}.
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#subnet_ids FsxOpenzfsFileSystem#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#throughput_capacity FsxOpenzfsFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `field:"required" json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#automatic_backup_retention_days FsxOpenzfsFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#backup_id FsxOpenzfsFileSystem#backup_id}.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#copy_tags_to_backups FsxOpenzfsFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#copy_tags_to_volumes FsxOpenzfsFileSystem#copy_tags_to_volumes}.
	CopyTagsToVolumes interface{} `field:"optional" json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#daily_automatic_backup_start_time FsxOpenzfsFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#delete_options FsxOpenzfsFileSystem#delete_options}.
	DeleteOptions *[]*string `field:"optional" json:"deleteOptions" yaml:"deleteOptions"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#disk_iops_configuration FsxOpenzfsFileSystem#disk_iops_configuration}
	DiskIopsConfiguration *FsxOpenzfsFileSystemDiskIopsConfiguration `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#endpoint_ip_address_range FsxOpenzfsFileSystem#endpoint_ip_address_range}.
	EndpointIpAddressRange *string `field:"optional" json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#final_backup_tags FsxOpenzfsFileSystem#final_backup_tags}.
	FinalBackupTags *map[string]*string `field:"optional" json:"finalBackupTags" yaml:"finalBackupTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#id FsxOpenzfsFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#kms_key_id FsxOpenzfsFileSystem#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#preferred_subnet_id FsxOpenzfsFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `field:"optional" json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#region FsxOpenzfsFileSystem#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// root_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#root_volume_configuration FsxOpenzfsFileSystem#root_volume_configuration}
	RootVolumeConfiguration *FsxOpenzfsFileSystemRootVolumeConfiguration `field:"optional" json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#route_table_ids FsxOpenzfsFileSystem#route_table_ids}.
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#security_group_ids FsxOpenzfsFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#skip_final_backup FsxOpenzfsFileSystem#skip_final_backup}.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#storage_capacity FsxOpenzfsFileSystem#storage_capacity}.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#storage_type FsxOpenzfsFileSystem#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#tags FsxOpenzfsFileSystem#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#tags_all FsxOpenzfsFileSystem#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#timeouts FsxOpenzfsFileSystem#timeouts}
	Timeouts *FsxOpenzfsFileSystemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#weekly_maintenance_start_time FsxOpenzfsFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

