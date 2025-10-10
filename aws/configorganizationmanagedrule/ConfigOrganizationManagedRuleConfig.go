// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configorganizationmanagedrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigOrganizationManagedRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#name ConfigOrganizationManagedRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#rule_identifier ConfigOrganizationManagedRule#rule_identifier}.
	RuleIdentifier *string `field:"required" json:"ruleIdentifier" yaml:"ruleIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#description ConfigOrganizationManagedRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#excluded_accounts ConfigOrganizationManagedRule#excluded_accounts}.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#id ConfigOrganizationManagedRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#input_parameters ConfigOrganizationManagedRule#input_parameters}.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#maximum_execution_frequency ConfigOrganizationManagedRule#maximum_execution_frequency}.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#region ConfigOrganizationManagedRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#resource_id_scope ConfigOrganizationManagedRule#resource_id_scope}.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#resource_types_scope ConfigOrganizationManagedRule#resource_types_scope}.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#tag_key_scope ConfigOrganizationManagedRule#tag_key_scope}.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#tag_value_scope ConfigOrganizationManagedRule#tag_value_scope}.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/config_organization_managed_rule#timeouts ConfigOrganizationManagedRule#timeouts}
	Timeouts *ConfigOrganizationManagedRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

