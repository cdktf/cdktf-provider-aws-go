// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxfilecache


type FsxFileCacheLustreConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/fsx_file_cache#deployment_type FsxFileCache#deployment_type}.
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/fsx_file_cache#metadata_configuration FsxFileCache#metadata_configuration}
	MetadataConfiguration interface{} `field:"required" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/fsx_file_cache#per_unit_storage_throughput FsxFileCache#per_unit_storage_throughput}.
	PerUnitStorageThroughput *float64 `field:"required" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/fsx_file_cache#weekly_maintenance_start_time FsxFileCache#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

