package servicecatalog


type ServicecatalogConstraintTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#create ServicecatalogConstraint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#delete ServicecatalogConstraint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#read ServicecatalogConstraint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/servicecatalog_constraint#update ServicecatalogConstraint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

