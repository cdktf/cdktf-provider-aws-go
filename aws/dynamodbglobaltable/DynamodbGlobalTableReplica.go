// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplica struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dynamodb_global_table#region_name DynamodbGlobalTable#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

