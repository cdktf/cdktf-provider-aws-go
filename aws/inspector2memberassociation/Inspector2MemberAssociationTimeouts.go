// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2memberassociation


type Inspector2MemberAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/inspector2_member_association#create Inspector2MemberAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/inspector2_member_association#delete Inspector2MemberAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

