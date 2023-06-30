package opensearchoutboundconnection


type OpensearchOutboundConnectionRemoteDomainInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/opensearch_outbound_connection#domain_name OpensearchOutboundConnection#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/opensearch_outbound_connection#owner_id OpensearchOutboundConnection#owner_id}.
	OwnerId *string `field:"required" json:"ownerId" yaml:"ownerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/opensearch_outbound_connection#region OpensearchOutboundConnection#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

