package servicecatalogproduct


type ServicecatalogProductTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/servicecatalog_product#create ServicecatalogProduct#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/servicecatalog_product#delete ServicecatalogProduct#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/servicecatalog_product#read ServicecatalogProduct#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/servicecatalog_product#update ServicecatalogProduct#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

