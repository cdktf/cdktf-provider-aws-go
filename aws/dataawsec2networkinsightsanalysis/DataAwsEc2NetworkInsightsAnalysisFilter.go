// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2networkinsightsanalysis


type DataAwsEc2NetworkInsightsAnalysisFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/ec2_network_insights_analysis#name DataAwsEc2NetworkInsightsAnalysis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/data-sources/ec2_network_insights_analysis#values DataAwsEc2NetworkInsightsAnalysis#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

