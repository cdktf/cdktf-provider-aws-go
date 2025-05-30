// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubConfigurationPolicyAssociationConfig struct {
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
	// The universally unique identifier (UUID) of the configuration policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securityhub_configuration_policy_association#policy_id SecurityhubConfigurationPolicyAssociation#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// The identifier of the target account, organizational unit, or the root to associate with the specified configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securityhub_configuration_policy_association#target_id SecurityhubConfigurationPolicyAssociation#target_id}
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securityhub_configuration_policy_association#id SecurityhubConfigurationPolicyAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securityhub_configuration_policy_association#timeouts SecurityhubConfigurationPolicyAssociation#timeouts}
	Timeouts *SecurityhubConfigurationPolicyAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

