// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteriaPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/inspector2_filter#begin_inclusive Inspector2Filter#begin_inclusive}.
	BeginInclusive *float64 `field:"required" json:"beginInclusive" yaml:"beginInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/inspector2_filter#end_inclusive Inspector2Filter#end_inclusive}.
	EndInclusive *float64 `field:"required" json:"endInclusive" yaml:"endInclusive"`
}

