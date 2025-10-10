// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapfilesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOntapFileSystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#deployment_type FsxOntapFileSystem#deployment_type}.
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#preferred_subnet_id FsxOntapFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `field:"required" json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#storage_capacity FsxOntapFileSystem#storage_capacity}.
	StorageCapacity *float64 `field:"required" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#subnet_ids FsxOntapFileSystem#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#automatic_backup_retention_days FsxOntapFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#daily_automatic_backup_start_time FsxOntapFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#disk_iops_configuration FsxOntapFileSystem#disk_iops_configuration}
	DiskIopsConfiguration *FsxOntapFileSystemDiskIopsConfiguration `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#endpoint_ip_address_range FsxOntapFileSystem#endpoint_ip_address_range}.
	EndpointIpAddressRange *string `field:"optional" json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#fsx_admin_password FsxOntapFileSystem#fsx_admin_password}.
	FsxAdminPassword *string `field:"optional" json:"fsxAdminPassword" yaml:"fsxAdminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#ha_pairs FsxOntapFileSystem#ha_pairs}.
	HaPairs *float64 `field:"optional" json:"haPairs" yaml:"haPairs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#id FsxOntapFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#kms_key_id FsxOntapFileSystem#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#region FsxOntapFileSystem#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#route_table_ids FsxOntapFileSystem#route_table_ids}.
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#security_group_ids FsxOntapFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#storage_type FsxOntapFileSystem#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#tags FsxOntapFileSystem#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#tags_all FsxOntapFileSystem#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#throughput_capacity FsxOntapFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#throughput_capacity_per_ha_pair FsxOntapFileSystem#throughput_capacity_per_ha_pair}.
	ThroughputCapacityPerHaPair *float64 `field:"optional" json:"throughputCapacityPerHaPair" yaml:"throughputCapacityPerHaPair"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#timeouts FsxOntapFileSystem#timeouts}
	Timeouts *FsxOntapFileSystemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/fsx_ontap_file_system#weekly_maintenance_start_time FsxOntapFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

