// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#acceleration_status S3Bucket#acceleration_status}.
	AccelerationStatus *string `field:"optional" json:"accelerationStatus" yaml:"accelerationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#acl S3Bucket#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#bucket S3Bucket#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#bucket_prefix S3Bucket#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#cors_rule S3Bucket#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#force_destroy S3Bucket#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#grant S3Bucket#grant}
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#id S3Bucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lifecycle_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#lifecycle_rule S3Bucket#lifecycle_rule}
	LifecycleRule interface{} `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#logging S3Bucket#logging}
	Logging *S3BucketLogging `field:"optional" json:"logging" yaml:"logging"`
	// object_lock_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#object_lock_configuration S3Bucket#object_lock_configuration}
	ObjectLockConfiguration *S3BucketObjectLockConfiguration `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}.
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#policy S3Bucket#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// replication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#replication_configuration S3Bucket#replication_configuration}
	ReplicationConfiguration *S3BucketReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#request_payer S3Bucket#request_payer}.
	RequestPayer *string `field:"optional" json:"requestPayer" yaml:"requestPayer"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#server_side_encryption_configuration S3Bucket#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *S3BucketServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#tags S3Bucket#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#tags_all S3Bucket#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#timeouts S3Bucket#timeouts}
	Timeouts *S3BucketTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// versioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#versioning S3Bucket#versioning}
	Versioning *S3BucketVersioning `field:"optional" json:"versioning" yaml:"versioning"`
	// website block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/s3_bucket#website S3Bucket#website}
	Website *S3BucketWebsite `field:"optional" json:"website" yaml:"website"`
}

