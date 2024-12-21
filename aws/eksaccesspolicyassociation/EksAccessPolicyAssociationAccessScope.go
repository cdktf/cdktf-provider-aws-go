// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksaccesspolicyassociation


type EksAccessPolicyAssociationAccessScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/eks_access_policy_association#type EksAccessPolicyAssociation#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/eks_access_policy_association#namespaces EksAccessPolicyAssociation#namespaces}.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

