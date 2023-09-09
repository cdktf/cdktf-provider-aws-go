// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationefs

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationEfsConfig struct {
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
	// ec2_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#ec2_config DatasyncLocationEfs#ec2_config}
	Ec2Config *DatasyncLocationEfsEc2Config `field:"required" json:"ec2Config" yaml:"ec2Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#efs_file_system_arn DatasyncLocationEfs#efs_file_system_arn}.
	EfsFileSystemArn *string `field:"required" json:"efsFileSystemArn" yaml:"efsFileSystemArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#access_point_arn DatasyncLocationEfs#access_point_arn}.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#file_system_access_role_arn DatasyncLocationEfs#file_system_access_role_arn}.
	FileSystemAccessRoleArn *string `field:"optional" json:"fileSystemAccessRoleArn" yaml:"fileSystemAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#id DatasyncLocationEfs#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#in_transit_encryption DatasyncLocationEfs#in_transit_encryption}.
	InTransitEncryption *string `field:"optional" json:"inTransitEncryption" yaml:"inTransitEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#subdirectory DatasyncLocationEfs#subdirectory}.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#tags DatasyncLocationEfs#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/datasync_location_efs#tags_all DatasyncLocationEfs#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

