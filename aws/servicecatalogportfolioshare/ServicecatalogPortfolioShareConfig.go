// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecatalogportfolioshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicecatalogPortfolioShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#portfolio_id ServicecatalogPortfolioShare#portfolio_id}.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#principal_id ServicecatalogPortfolioShare#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#type ServicecatalogPortfolioShare#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#accept_language ServicecatalogPortfolioShare#accept_language}.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#id ServicecatalogPortfolioShare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#share_principals ServicecatalogPortfolioShare#share_principals}.
	SharePrincipals interface{} `field:"optional" json:"sharePrincipals" yaml:"sharePrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#share_tag_options ServicecatalogPortfolioShare#share_tag_options}.
	ShareTagOptions interface{} `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#timeouts ServicecatalogPortfolioShare#timeouts}
	Timeouts *ServicecatalogPortfolioShareTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/servicecatalog_portfolio_share#wait_for_acceptance ServicecatalogPortfolioShare#wait_for_acceptance}.
	WaitForAcceptance interface{} `field:"optional" json:"waitForAcceptance" yaml:"waitForAcceptance"`
}

