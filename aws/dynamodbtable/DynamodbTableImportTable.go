// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableImportTable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dynamodb_table#input_format DynamodbTable#input_format}.
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// s3_bucket_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dynamodb_table#s3_bucket_source DynamodbTable#s3_bucket_source}
	S3BucketSource *DynamodbTableImportTableS3BucketSource `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dynamodb_table#input_compression_type DynamodbTable#input_compression_type}.
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// input_format_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/dynamodb_table#input_format_options DynamodbTable#input_format_options}
	InputFormatOptions *DynamodbTableImportTableInputFormatOptions `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

