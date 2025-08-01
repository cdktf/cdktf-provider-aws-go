// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerlinkassociation


type NetworkmanagerLinkAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkmanager_link_association#create NetworkmanagerLinkAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkmanager_link_association#delete NetworkmanagerLinkAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

