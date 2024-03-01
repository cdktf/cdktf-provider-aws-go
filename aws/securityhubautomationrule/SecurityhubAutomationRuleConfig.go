// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubautomationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubAutomationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#description SecurityhubAutomationRule#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#rule_name SecurityhubAutomationRule#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#rule_order SecurityhubAutomationRule#rule_order}.
	RuleOrder *float64 `field:"required" json:"ruleOrder" yaml:"ruleOrder"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#actions SecurityhubAutomationRule#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#criteria SecurityhubAutomationRule#criteria}
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#is_terminal SecurityhubAutomationRule#is_terminal}.
	IsTerminal interface{} `field:"optional" json:"isTerminal" yaml:"isTerminal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#rule_status SecurityhubAutomationRule#rule_status}.
	RuleStatus *string `field:"optional" json:"ruleStatus" yaml:"ruleStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/securityhub_automation_rule#tags SecurityhubAutomationRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

