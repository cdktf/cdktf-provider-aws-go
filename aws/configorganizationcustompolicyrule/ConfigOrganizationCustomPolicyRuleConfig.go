// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configorganizationcustompolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigOrganizationCustomPolicyRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#name ConfigOrganizationCustomPolicyRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#policy_runtime ConfigOrganizationCustomPolicyRule#policy_runtime}.
	PolicyRuntime *string `field:"required" json:"policyRuntime" yaml:"policyRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#policy_text ConfigOrganizationCustomPolicyRule#policy_text}.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#trigger_types ConfigOrganizationCustomPolicyRule#trigger_types}.
	TriggerTypes *[]*string `field:"required" json:"triggerTypes" yaml:"triggerTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#debug_log_delivery_accounts ConfigOrganizationCustomPolicyRule#debug_log_delivery_accounts}.
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#description ConfigOrganizationCustomPolicyRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#excluded_accounts ConfigOrganizationCustomPolicyRule#excluded_accounts}.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#id ConfigOrganizationCustomPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#input_parameters ConfigOrganizationCustomPolicyRule#input_parameters}.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#maximum_execution_frequency ConfigOrganizationCustomPolicyRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#region ConfigOrganizationCustomPolicyRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#resource_id_scope ConfigOrganizationCustomPolicyRule#resource_id_scope}.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#resource_types_scope ConfigOrganizationCustomPolicyRule#resource_types_scope}.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#tag_key_scope ConfigOrganizationCustomPolicyRule#tag_key_scope}.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#tag_value_scope ConfigOrganizationCustomPolicyRule#tag_value_scope}.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_organization_custom_policy_rule#timeouts ConfigOrganizationCustomPolicyRule#timeouts}
	Timeouts *ConfigOrganizationCustomPolicyRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

