// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bcmdataexports_export#name BcmdataexportsExport#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// data_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bcmdataexports_export#data_query BcmdataexportsExport#data_query}
	DataQuery interface{} `field:"optional" json:"dataQuery" yaml:"dataQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bcmdataexports_export#description BcmdataexportsExport#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bcmdataexports_export#destination_configurations BcmdataexportsExport#destination_configurations}
	DestinationConfigurations interface{} `field:"optional" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// refresh_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bcmdataexports_export#refresh_cadence BcmdataexportsExport#refresh_cadence}
	RefreshCadence interface{} `field:"optional" json:"refreshCadence" yaml:"refreshCadence"`
}

