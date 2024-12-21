// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExportDestinationConfigurationsS3Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/bcmdataexports_export#s3_bucket BcmdataexportsExport#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/bcmdataexports_export#s3_prefix BcmdataexportsExport#s3_prefix}.
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/bcmdataexports_export#s3_region BcmdataexportsExport#s3_region}.
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// s3_output_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/bcmdataexports_export#s3_output_configurations BcmdataexportsExport#s3_output_configurations}
	S3OutputConfigurations interface{} `field:"optional" json:"s3OutputConfigurations" yaml:"s3OutputConfigurations"`
}

