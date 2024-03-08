// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalratebasedrule


type WafregionalRateBasedRulePredicate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rate_based_rule#data_id WafregionalRateBasedRule#data_id}.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rate_based_rule#negated WafregionalRateBasedRule#negated}.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rate_based_rule#type WafregionalRateBasedRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

