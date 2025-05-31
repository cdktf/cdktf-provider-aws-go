// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bcmdataexports_export#compression BcmdataexportsExport#compression}.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bcmdataexports_export#format BcmdataexportsExport#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bcmdataexports_export#output_type BcmdataexportsExport#output_type}.
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bcmdataexports_export#overwrite BcmdataexportsExport#overwrite}.
	Overwrite *string `field:"required" json:"overwrite" yaml:"overwrite"`
}

