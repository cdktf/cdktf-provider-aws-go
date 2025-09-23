// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersUpdatedAt struct {
	// date_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/securityhub_insight#date_range SecurityhubInsight#date_range}
	DateRange *SecurityhubInsightFiltersUpdatedAtDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/securityhub_insight#end SecurityhubInsight#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/securityhub_insight#start SecurityhubInsight#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

