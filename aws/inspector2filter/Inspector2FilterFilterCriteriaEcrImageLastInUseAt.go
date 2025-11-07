// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteriaEcrImageLastInUseAt struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/inspector2_filter#end_inclusive Inspector2Filter#end_inclusive}.
	EndInclusive *string `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/inspector2_filter#start_inclusive Inspector2Filter#start_inclusive}.
	StartInclusive *string `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

