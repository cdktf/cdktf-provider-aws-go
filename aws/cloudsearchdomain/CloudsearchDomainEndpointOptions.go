package cloudsearchdomain


type CloudsearchDomainEndpointOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudsearch_domain#enforce_https CloudsearchDomain#enforce_https}.
	EnforceHttps interface{} `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudsearch_domain#tls_security_policy CloudsearchDomain#tls_security_policy}.
	TlsSecurityPolicy *string `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

