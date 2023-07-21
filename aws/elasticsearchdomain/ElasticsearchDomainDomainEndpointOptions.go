package elasticsearchdomain


type ElasticsearchDomainDomainEndpointOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/elasticsearch_domain#custom_endpoint ElasticsearchDomain#custom_endpoint}.
	CustomEndpoint *string `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/elasticsearch_domain#custom_endpoint_certificate_arn ElasticsearchDomain#custom_endpoint_certificate_arn}.
	CustomEndpointCertificateArn *string `field:"optional" json:"customEndpointCertificateArn" yaml:"customEndpointCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/elasticsearch_domain#custom_endpoint_enabled ElasticsearchDomain#custom_endpoint_enabled}.
	CustomEndpointEnabled interface{} `field:"optional" json:"customEndpointEnabled" yaml:"customEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/elasticsearch_domain#enforce_https ElasticsearchDomain#enforce_https}.
	EnforceHttps interface{} `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/elasticsearch_domain#tls_security_policy ElasticsearchDomain#tls_security_policy}.
	TlsSecurityPolicy *string `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

