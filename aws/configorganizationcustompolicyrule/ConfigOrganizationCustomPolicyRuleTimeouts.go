// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configorganizationcustompolicyrule


type ConfigOrganizationCustomPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/config_organization_custom_policy_rule#create ConfigOrganizationCustomPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/config_organization_custom_policy_rule#delete ConfigOrganizationCustomPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/config_organization_custom_policy_rule#update ConfigOrganizationCustomPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

