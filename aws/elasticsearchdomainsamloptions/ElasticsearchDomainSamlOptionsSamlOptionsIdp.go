package elasticsearchdomainsamloptions


type ElasticsearchDomainSamlOptionsSamlOptionsIdp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#entity_id ElasticsearchDomainSamlOptions#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#metadata_content ElasticsearchDomainSamlOptions#metadata_content}.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

