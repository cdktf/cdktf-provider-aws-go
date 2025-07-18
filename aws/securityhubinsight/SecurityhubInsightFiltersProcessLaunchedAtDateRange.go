// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersProcessLaunchedAtDateRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/securityhub_insight#unit SecurityhubInsight#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/securityhub_insight#value SecurityhubInsight#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

