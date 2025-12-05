// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomainsamloptions


type OpensearchDomainSamlOptionsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/opensearch_domain_saml_options#delete OpensearchDomainSamlOptions#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/opensearch_domain_saml_options#update OpensearchDomainSamlOptions#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

