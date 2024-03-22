// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersNetworkSourcePort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/securityhub_insight#eq SecurityhubInsight#eq}.
	Eq *string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/securityhub_insight#gte SecurityhubInsight#gte}.
	Gte *string `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/securityhub_insight#lte SecurityhubInsight#lte}.
	Lte *string `field:"optional" json:"lte" yaml:"lte"`
}

