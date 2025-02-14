// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsefsmounttarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEfsMountTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/efs_mount_target#access_point_id DataAwsEfsMountTarget#access_point_id}.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/efs_mount_target#file_system_id DataAwsEfsMountTarget#file_system_id}.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/efs_mount_target#id DataAwsEfsMountTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/efs_mount_target#mount_target_id DataAwsEfsMountTarget#mount_target_id}.
	MountTargetId *string `field:"optional" json:"mountTargetId" yaml:"mountTargetId"`
}

