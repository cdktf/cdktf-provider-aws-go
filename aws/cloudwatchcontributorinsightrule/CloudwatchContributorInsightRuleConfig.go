// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchcontributorinsightrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchContributorInsightRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_contributor_insight_rule#rule_definition CloudwatchContributorInsightRule#rule_definition}.
	RuleDefinition *string `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_contributor_insight_rule#rule_name CloudwatchContributorInsightRule#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_contributor_insight_rule#rule_state CloudwatchContributorInsightRule#rule_state}.
	RuleState *string `field:"optional" json:"ruleState" yaml:"ruleState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_contributor_insight_rule#tags CloudwatchContributorInsightRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

