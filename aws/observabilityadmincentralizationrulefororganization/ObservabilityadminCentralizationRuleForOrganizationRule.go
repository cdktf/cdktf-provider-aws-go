// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRule struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination ObservabilityadminCentralizationRuleForOrganization#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/observabilityadmin_centralization_rule_for_organization#source ObservabilityadminCentralizationRuleForOrganization#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

