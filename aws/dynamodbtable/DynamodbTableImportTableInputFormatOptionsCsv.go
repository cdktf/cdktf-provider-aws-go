// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableImportTableInputFormatOptionsCsv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/dynamodb_table#delimiter DynamodbTable#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/dynamodb_table#header_list DynamodbTable#header_list}.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

