// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersNetworkDestinationIpv6 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/securityhub_insight#cidr SecurityhubInsight#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

