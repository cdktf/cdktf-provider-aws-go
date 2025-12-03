// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingview


type BillingViewDataFilterExpressionTimeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/billing_view#begin_date_inclusive BillingView#begin_date_inclusive}.
	BeginDateInclusive *string `field:"required" json:"beginDateInclusive" yaml:"beginDateInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/billing_view#end_date_inclusive BillingView#end_date_inclusive}.
	EndDateInclusive *string `field:"required" json:"endDateInclusive" yaml:"endDateInclusive"`
}

