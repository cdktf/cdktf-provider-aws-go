// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtableexport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DynamodbTableExportConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#s3_bucket DynamodbTableExport#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#table_arn DynamodbTableExport#table_arn}.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#export_format DynamodbTableExport#export_format}.
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#export_time DynamodbTableExport#export_time}.
	ExportTime *string `field:"optional" json:"exportTime" yaml:"exportTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#export_type DynamodbTableExport#export_type}.
	ExportType *string `field:"optional" json:"exportType" yaml:"exportType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#id DynamodbTableExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incremental_export_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#incremental_export_specification DynamodbTableExport#incremental_export_specification}
	IncrementalExportSpecification *DynamodbTableExportIncrementalExportSpecification `field:"optional" json:"incrementalExportSpecification" yaml:"incrementalExportSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#s3_bucket_owner DynamodbTableExport#s3_bucket_owner}.
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#s3_prefix DynamodbTableExport#s3_prefix}.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#s3_sse_algorithm DynamodbTableExport#s3_sse_algorithm}.
	S3SseAlgorithm *string `field:"optional" json:"s3SseAlgorithm" yaml:"s3SseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#s3_sse_kms_key_id DynamodbTableExport#s3_sse_kms_key_id}.
	S3SseKmsKeyId *string `field:"optional" json:"s3SseKmsKeyId" yaml:"s3SseKmsKeyId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/dynamodb_table_export#timeouts DynamodbTableExport#timeouts}
	Timeouts *DynamodbTableExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

