// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard


type QuicksightDashboardParametersDecimalParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

