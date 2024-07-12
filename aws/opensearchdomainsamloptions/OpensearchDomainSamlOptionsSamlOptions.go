// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomainsamloptions


type OpensearchDomainSamlOptionsSamlOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#enabled OpensearchDomainSamlOptions#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#idp OpensearchDomainSamlOptions#idp}
	Idp *OpensearchDomainSamlOptionsSamlOptionsIdp `field:"optional" json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#master_backend_role OpensearchDomainSamlOptions#master_backend_role}.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#master_user_name OpensearchDomainSamlOptions#master_user_name}.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#roles_key OpensearchDomainSamlOptions#roles_key}.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#session_timeout_minutes OpensearchDomainSamlOptions#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/opensearch_domain_saml_options#subject_key OpensearchDomainSamlOptions#subject_key}.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

