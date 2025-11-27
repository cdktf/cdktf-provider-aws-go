// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingview


type BillingViewDataFilterExpressionDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#key BillingView#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#values BillingView#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

