// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslaunchtemplate


type DataAwsLaunchTemplateFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/data-sources/launch_template#name DataAwsLaunchTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/data-sources/launch_template#values DataAwsLaunchTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

