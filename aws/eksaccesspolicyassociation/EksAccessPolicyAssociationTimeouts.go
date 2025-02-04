// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksaccesspolicyassociation


type EksAccessPolicyAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/eks_access_policy_association#create EksAccessPolicyAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/eks_access_policy_association#delete EksAccessPolicyAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

