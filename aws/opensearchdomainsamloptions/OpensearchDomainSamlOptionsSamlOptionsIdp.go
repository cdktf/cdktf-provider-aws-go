package opensearchdomainsamloptions


type OpensearchDomainSamlOptionsSamlOptionsIdp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/opensearch_domain_saml_options#entity_id OpensearchDomainSamlOptions#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/opensearch_domain_saml_options#metadata_content OpensearchDomainSamlOptions#metadata_content}.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

