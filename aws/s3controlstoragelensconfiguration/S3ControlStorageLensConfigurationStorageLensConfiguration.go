// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfiguration struct {
	// account_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#account_level S3ControlStorageLensConfiguration#account_level}
	AccountLevel *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#enabled S3ControlStorageLensConfiguration#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// aws_org block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#aws_org S3ControlStorageLensConfiguration#aws_org}
	AwsOrg *S3ControlStorageLensConfigurationStorageLensConfigurationAwsOrg `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#data_export S3ControlStorageLensConfiguration#data_export}
	DataExport *S3ControlStorageLensConfigurationStorageLensConfigurationDataExport `field:"optional" json:"dataExport" yaml:"dataExport"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#exclude S3ControlStorageLensConfiguration#exclude}
	Exclude *S3ControlStorageLensConfigurationStorageLensConfigurationExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/s3control_storage_lens_configuration#include S3ControlStorageLensConfiguration#include}
	Include *S3ControlStorageLensConfigurationStorageLensConfigurationInclude `field:"optional" json:"include" yaml:"include"`
}

