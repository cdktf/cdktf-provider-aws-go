// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/imagebuilder_distribution_configuration#launch_template_id ImagebuilderDistributionConfiguration#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/imagebuilder_distribution_configuration#launch_template_name ImagebuilderDistributionConfiguration#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/imagebuilder_distribution_configuration#launch_template_version ImagebuilderDistributionConfiguration#launch_template_version}.
	LaunchTemplateVersion *string `field:"optional" json:"launchTemplateVersion" yaml:"launchTemplateVersion"`
}

