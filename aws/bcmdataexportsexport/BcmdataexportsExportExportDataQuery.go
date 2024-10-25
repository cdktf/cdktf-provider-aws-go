// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExportDataQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/bcmdataexports_export#query_statement BcmdataexportsExport#query_statement}.
	QueryStatement *string `field:"required" json:"queryStatement" yaml:"queryStatement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/bcmdataexports_export#table_configurations BcmdataexportsExport#table_configurations}.
	TableConfigurations interface{} `field:"optional" json:"tableConfigurations" yaml:"tableConfigurations"`
}

