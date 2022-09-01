package servicecatalog


type ServicecatalogPrincipalPortfolioAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#create ServicecatalogPrincipalPortfolioAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#delete ServicecatalogPrincipalPortfolioAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_principal_portfolio_association#read ServicecatalogPrincipalPortfolioAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

