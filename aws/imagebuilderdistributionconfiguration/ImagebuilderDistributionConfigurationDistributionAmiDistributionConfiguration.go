// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#ami_tags ImagebuilderDistributionConfiguration#ami_tags}.
	AmiTags *map[string]*string `field:"optional" json:"amiTags" yaml:"amiTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#description ImagebuilderDistributionConfiguration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#kms_key_id ImagebuilderDistributionConfiguration#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// launch_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#launch_permission ImagebuilderDistributionConfiguration#launch_permission}
	LaunchPermission *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermission `field:"optional" json:"launchPermission" yaml:"launchPermission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#name ImagebuilderDistributionConfiguration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/imagebuilder_distribution_configuration#target_account_ids ImagebuilderDistributionConfiguration#target_account_ids}.
	TargetAccountIds *[]*string `field:"optional" json:"targetAccountIds" yaml:"targetAccountIds"`
}

