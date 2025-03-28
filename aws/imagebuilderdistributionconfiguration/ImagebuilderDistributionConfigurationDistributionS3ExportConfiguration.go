// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/imagebuilder_distribution_configuration#disk_image_format ImagebuilderDistributionConfiguration#disk_image_format}.
	DiskImageFormat *string `field:"required" json:"diskImageFormat" yaml:"diskImageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/imagebuilder_distribution_configuration#role_name ImagebuilderDistributionConfiguration#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/imagebuilder_distribution_configuration#s3_bucket ImagebuilderDistributionConfiguration#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/imagebuilder_distribution_configuration#s3_prefix ImagebuilderDistributionConfiguration#s3_prefix}.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
}

