package opensearchdomainpolicy


type OpensearchDomainPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain_policy#delete OpensearchDomainPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opensearch_domain_policy#update OpensearchDomainPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

