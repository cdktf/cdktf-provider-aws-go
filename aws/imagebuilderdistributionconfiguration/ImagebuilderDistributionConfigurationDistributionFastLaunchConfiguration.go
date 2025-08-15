// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionFastLaunchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/imagebuilder_distribution_configuration#account_id ImagebuilderDistributionConfiguration#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/imagebuilder_distribution_configuration#enabled ImagebuilderDistributionConfiguration#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/imagebuilder_distribution_configuration#launch_template ImagebuilderDistributionConfiguration#launch_template}
	LaunchTemplate *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/imagebuilder_distribution_configuration#max_parallel_launches ImagebuilderDistributionConfiguration#max_parallel_launches}.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/imagebuilder_distribution_configuration#snapshot_configuration ImagebuilderDistributionConfiguration#snapshot_configuration}
	SnapshotConfiguration *ImagebuilderDistributionConfigurationDistributionFastLaunchConfigurationSnapshotConfiguration `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

