package elasticacheuser


type ElasticacheUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticache_user#create ElasticacheUser#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticache_user#delete ElasticacheUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elasticache_user#update ElasticacheUser#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

