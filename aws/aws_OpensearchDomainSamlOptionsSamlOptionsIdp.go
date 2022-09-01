// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type OpensearchDomainSamlOptionsSamlOptionsIdp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain_saml_options#entity_id OpensearchDomainSamlOptions#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain_saml_options#metadata_content OpensearchDomainSamlOptions#metadata_content}.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

