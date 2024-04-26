// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketinventory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketInventoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#bucket S3BucketInventory#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#destination S3BucketInventory#destination}
	Destination *S3BucketInventoryDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#included_object_versions S3BucketInventory#included_object_versions}.
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#name S3BucketInventory#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#schedule S3BucketInventory#schedule}
	Schedule *S3BucketInventorySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#enabled S3BucketInventory#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#filter S3BucketInventory#filter}
	Filter *S3BucketInventoryFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#id S3BucketInventory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/s3_bucket_inventory#optional_fields S3BucketInventory#optional_fields}.
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
}

