// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogproductportfolioassociation


type ServicecatalogProductPortfolioAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalog_product_portfolio_association#create ServicecatalogProductPortfolioAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalog_product_portfolio_association#delete ServicecatalogProductPortfolioAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/servicecatalog_product_portfolio_association#read ServicecatalogProductPortfolioAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

