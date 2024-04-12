// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicyassociation


type SecurityhubConfigurationPolicyAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_configuration_policy_association#create SecurityhubConfigurationPolicyAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securityhub_configuration_policy_association#update SecurityhubConfigurationPolicyAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

