// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagercontrol


type AuditmanagerControlControlMappingSourcesSourceKeyword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/auditmanager_control#keyword_input_type AuditmanagerControl#keyword_input_type}.
	KeywordInputType *string `field:"optional" json:"keywordInputType" yaml:"keywordInputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/auditmanager_control#keyword_value AuditmanagerControl#keyword_value}.
	KeywordValue *string `field:"optional" json:"keywordValue" yaml:"keywordValue"`
}

