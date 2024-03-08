// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalrule


type WafregionalRulePredicate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rule#data_id WafregionalRule#data_id}.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rule#negated WafregionalRule#negated}.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_rule#type WafregionalRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

