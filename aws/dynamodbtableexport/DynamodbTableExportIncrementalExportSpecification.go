// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtableexport


type DynamodbTableExportIncrementalExportSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/dynamodb_table_export#export_from_time DynamodbTableExport#export_from_time}.
	ExportFromTime *string `field:"optional" json:"exportFromTime" yaml:"exportFromTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/dynamodb_table_export#export_to_time DynamodbTableExport#export_to_time}.
	ExportToTime *string `field:"optional" json:"exportToTime" yaml:"exportToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/dynamodb_table_export#export_view_type DynamodbTableExport#export_view_type}.
	ExportViewType *string `field:"optional" json:"exportViewType" yaml:"exportViewType"`
}

