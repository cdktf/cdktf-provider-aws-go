// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest


type SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#id SpotFleetRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#name SpotFleetRequest#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/spot_fleet_request#version SpotFleetRequest#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

