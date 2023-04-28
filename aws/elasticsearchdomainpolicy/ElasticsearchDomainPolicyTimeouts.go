package elasticsearchdomainpolicy


type ElasticsearchDomainPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/elasticsearch_domain_policy#delete ElasticsearchDomainPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/elasticsearch_domain_policy#update ElasticsearchDomainPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

