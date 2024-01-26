// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amplifydomainassociation


type AmplifyDomainAssociationSubDomain struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/amplify_domain_association#branch_name AmplifyDomainAssociation#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/amplify_domain_association#prefix AmplifyDomainAssociation#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

