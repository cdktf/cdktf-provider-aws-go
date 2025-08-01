// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistribution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#region ImagebuilderDistributionConfiguration#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// ami_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#ami_distribution_configuration ImagebuilderDistributionConfiguration#ami_distribution_configuration}
	AmiDistributionConfiguration *ImagebuilderDistributionConfigurationDistributionAmiDistributionConfiguration `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// container_distribution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#container_distribution_configuration ImagebuilderDistributionConfiguration#container_distribution_configuration}
	ContainerDistributionConfiguration *ImagebuilderDistributionConfigurationDistributionContainerDistributionConfiguration `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// fast_launch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#fast_launch_configuration ImagebuilderDistributionConfiguration#fast_launch_configuration}
	FastLaunchConfiguration interface{} `field:"optional" json:"fastLaunchConfiguration" yaml:"fastLaunchConfiguration"`
	// launch_template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#launch_template_configuration ImagebuilderDistributionConfiguration#launch_template_configuration}
	LaunchTemplateConfiguration interface{} `field:"optional" json:"launchTemplateConfiguration" yaml:"launchTemplateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#license_configuration_arns ImagebuilderDistributionConfiguration#license_configuration_arns}.
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// s3_export_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#s3_export_configuration ImagebuilderDistributionConfiguration#s3_export_configuration}
	S3ExportConfiguration *ImagebuilderDistributionConfigurationDistributionS3ExportConfiguration `field:"optional" json:"s3ExportConfiguration" yaml:"s3ExportConfiguration"`
	// ssm_parameter_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/imagebuilder_distribution_configuration#ssm_parameter_configuration ImagebuilderDistributionConfiguration#ssm_parameter_configuration}
	SsmParameterConfiguration interface{} `field:"optional" json:"ssmParameterConfiguration" yaml:"ssmParameterConfiguration"`
}

