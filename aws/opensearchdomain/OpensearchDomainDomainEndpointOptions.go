// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainDomainEndpointOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_domain#custom_endpoint OpensearchDomain#custom_endpoint}.
	CustomEndpoint *string `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_domain#custom_endpoint_certificate_arn OpensearchDomain#custom_endpoint_certificate_arn}.
	CustomEndpointCertificateArn *string `field:"optional" json:"customEndpointCertificateArn" yaml:"customEndpointCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_domain#custom_endpoint_enabled OpensearchDomain#custom_endpoint_enabled}.
	CustomEndpointEnabled interface{} `field:"optional" json:"customEndpointEnabled" yaml:"customEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_domain#enforce_https OpensearchDomain#enforce_https}.
	EnforceHttps interface{} `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/opensearch_domain#tls_security_policy OpensearchDomain#tls_security_policy}.
	TlsSecurityPolicy *string `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

