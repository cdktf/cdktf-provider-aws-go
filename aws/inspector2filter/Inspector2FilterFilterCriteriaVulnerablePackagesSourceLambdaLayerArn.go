// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteriaVulnerablePackagesSourceLambdaLayerArn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/inspector2_filter#comparison Inspector2Filter#comparison}.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/inspector2_filter#value Inspector2Filter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

