// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogtagoptionresourceassociation


type ServicecatalogTagOptionResourceAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_tag_option_resource_association#create ServicecatalogTagOptionResourceAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_tag_option_resource_association#delete ServicecatalogTagOptionResourceAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_tag_option_resource_association#read ServicecatalogTagOptionResourceAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

