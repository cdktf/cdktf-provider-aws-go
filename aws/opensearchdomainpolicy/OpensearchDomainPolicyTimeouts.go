package opensearchdomainpolicy


type OpensearchDomainPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/opensearch_domain_policy#delete OpensearchDomainPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/opensearch_domain_policy#update OpensearchDomainPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

