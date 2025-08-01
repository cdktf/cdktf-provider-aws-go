// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3tablestablebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3TablesTableBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table_bucket#name S3TablesTableBucket#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table_bucket#encryption_configuration S3TablesTableBucket#encryption_configuration}.
	EncryptionConfiguration *S3TablesTableBucketEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table_bucket#maintenance_configuration S3TablesTableBucket#maintenance_configuration}.
	MaintenanceConfiguration *S3TablesTableBucketMaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3tables_table_bucket#region S3TablesTableBucket#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

