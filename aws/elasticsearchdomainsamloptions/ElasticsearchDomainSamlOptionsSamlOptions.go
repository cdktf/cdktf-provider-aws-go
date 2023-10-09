// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomainsamloptions


type ElasticsearchDomainSamlOptionsSamlOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#enabled ElasticsearchDomainSamlOptions#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#idp ElasticsearchDomainSamlOptions#idp}
	Idp *ElasticsearchDomainSamlOptionsSamlOptionsIdp `field:"optional" json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#master_backend_role ElasticsearchDomainSamlOptions#master_backend_role}.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#master_user_name ElasticsearchDomainSamlOptions#master_user_name}.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#roles_key ElasticsearchDomainSamlOptions#roles_key}.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#session_timeout_minutes ElasticsearchDomainSamlOptions#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/elasticsearch_domain_saml_options#subject_key ElasticsearchDomainSamlOptions#subject_key}.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

