package elasticsearch


type ElasticsearchDomainPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy#delete ElasticsearchDomainPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy#update ElasticsearchDomainPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

