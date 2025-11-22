// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityadminCentralizationRuleForOrganizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#rule_name ObservabilityadminCentralizationRuleForOrganization#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#region ObservabilityadminCentralizationRuleForOrganization#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#rule ObservabilityadminCentralizationRuleForOrganization#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#tags ObservabilityadminCentralizationRuleForOrganization#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#timeouts ObservabilityadminCentralizationRuleForOrganization#timeouts}
	Timeouts *ObservabilityadminCentralizationRuleForOrganizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

