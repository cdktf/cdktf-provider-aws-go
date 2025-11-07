// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainIdentityCenterOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/opensearch_domain#enabled_api_access OpensearchDomain#enabled_api_access}.
	EnabledApiAccess interface{} `field:"optional" json:"enabledApiAccess" yaml:"enabledApiAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/opensearch_domain#identity_center_instance_arn OpensearchDomain#identity_center_instance_arn}.
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/opensearch_domain#roles_key OpensearchDomain#roles_key}.
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/opensearch_domain#subject_key OpensearchDomain#subject_key}.
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

