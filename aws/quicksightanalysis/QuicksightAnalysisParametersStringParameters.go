// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis


type QuicksightAnalysisParametersStringParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_analysis#values QuicksightAnalysis#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

