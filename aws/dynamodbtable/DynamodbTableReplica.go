// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableReplica struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/dynamodb_table#region_name DynamodbTable#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/dynamodb_table#kms_key_arn DynamodbTable#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/dynamodb_table#point_in_time_recovery DynamodbTable#point_in_time_recovery}.
	PointInTimeRecovery interface{} `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/dynamodb_table#propagate_tags DynamodbTable#propagate_tags}.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

