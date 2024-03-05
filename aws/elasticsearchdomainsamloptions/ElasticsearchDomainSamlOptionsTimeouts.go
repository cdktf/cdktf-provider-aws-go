// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomainsamloptions


type ElasticsearchDomainSamlOptionsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/elasticsearch_domain_saml_options#delete ElasticsearchDomainSamlOptions#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/elasticsearch_domain_saml_options#update ElasticsearchDomainSamlOptions#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

