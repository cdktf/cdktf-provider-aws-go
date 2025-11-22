// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingview


type BillingViewDataFilterExpression struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/billing_view#dimensions BillingView#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/billing_view#tags BillingView#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// time_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/billing_view#time_range BillingView#time_range}
	TimeRange interface{} `field:"optional" json:"timeRange" yaml:"timeRange"`
}

