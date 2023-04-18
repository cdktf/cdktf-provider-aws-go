package elasticsearchdomain


type ElasticsearchDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/elasticsearch_domain#create ElasticsearchDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/elasticsearch_domain#delete ElasticsearchDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/elasticsearch_domain#update ElasticsearchDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

