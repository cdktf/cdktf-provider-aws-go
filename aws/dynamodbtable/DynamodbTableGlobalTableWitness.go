// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableGlobalTableWitness struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/dynamodb_table#region_name DynamodbTable#region_name}.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

