// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteriaVulnerablePackagesEpoch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/inspector2_filter#lower_inclusive Inspector2Filter#lower_inclusive}.
	LowerInclusive *float64 `field:"required" json:"lowerInclusive" yaml:"lowerInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/inspector2_filter#upper_inclusive Inspector2Filter#upper_inclusive}.
	UpperInclusive *float64 `field:"required" json:"upperInclusive" yaml:"upperInclusive"`
}

