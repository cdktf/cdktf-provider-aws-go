// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtableexport


type DynamodbTableExportTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/dynamodb_table_export#create DynamodbTableExport#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/dynamodb_table_export#delete DynamodbTableExport#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

