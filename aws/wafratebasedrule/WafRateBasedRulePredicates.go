// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafratebasedrule


type WafRateBasedRulePredicates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/waf_rate_based_rule#data_id WafRateBasedRule#data_id}.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/waf_rate_based_rule#negated WafRateBasedRule#negated}.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/waf_rate_based_rule#type WafRateBasedRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

