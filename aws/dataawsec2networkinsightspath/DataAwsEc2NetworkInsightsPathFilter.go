// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2networkinsightspath


type DataAwsEc2NetworkInsightsPathFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/ec2_network_insights_path#name DataAwsEc2NetworkInsightsPath#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/ec2_network_insights_path#values DataAwsEc2NetworkInsightsPath#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

