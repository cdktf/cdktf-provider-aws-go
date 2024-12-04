// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest


type SpotFleetRequestLaunchTemplateConfig struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/spot_fleet_request#launch_template_specification SpotFleetRequest#launch_template_specification}
	LaunchTemplateSpecification *SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecification `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/spot_fleet_request#overrides SpotFleetRequest#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

