// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configremediationconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigRemediationConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#config_rule_name ConfigRemediationConfiguration#config_rule_name}.
	ConfigRuleName *string `field:"required" json:"configRuleName" yaml:"configRuleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#target_id ConfigRemediationConfiguration#target_id}.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#target_type ConfigRemediationConfiguration#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#automatic ConfigRemediationConfiguration#automatic}.
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// execution_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#execution_controls ConfigRemediationConfiguration#execution_controls}
	ExecutionControls *ConfigRemediationConfigurationExecutionControls `field:"optional" json:"executionControls" yaml:"executionControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#id ConfigRemediationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#maximum_automatic_attempts ConfigRemediationConfiguration#maximum_automatic_attempts}.
	MaximumAutomaticAttempts *float64 `field:"optional" json:"maximumAutomaticAttempts" yaml:"maximumAutomaticAttempts"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#parameter ConfigRemediationConfiguration#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#resource_type ConfigRemediationConfiguration#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#retry_attempt_seconds ConfigRemediationConfiguration#retry_attempt_seconds}.
	RetryAttemptSeconds *float64 `field:"optional" json:"retryAttemptSeconds" yaml:"retryAttemptSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/config_remediation_configuration#target_version ConfigRemediationConfiguration#target_version}.
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

