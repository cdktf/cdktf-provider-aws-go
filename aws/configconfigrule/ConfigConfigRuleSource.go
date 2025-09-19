// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigrule


type ConfigConfigRuleSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/config_config_rule#owner ConfigConfigRule#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// custom_policy_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/config_config_rule#custom_policy_details ConfigConfigRule#custom_policy_details}
	CustomPolicyDetails *ConfigConfigRuleSourceCustomPolicyDetails `field:"optional" json:"customPolicyDetails" yaml:"customPolicyDetails"`
	// source_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/config_config_rule#source_detail ConfigConfigRule#source_detail}
	SourceDetail interface{} `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/config_config_rule#source_identifier ConfigConfigRule#source_identifier}.
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

