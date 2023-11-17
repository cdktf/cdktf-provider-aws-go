// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogportfolio


type ServicecatalogPortfolioTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/servicecatalog_portfolio#create ServicecatalogPortfolio#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/servicecatalog_portfolio#delete ServicecatalogPortfolio#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/servicecatalog_portfolio#read ServicecatalogPortfolio#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/servicecatalog_portfolio#update ServicecatalogPortfolio#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

