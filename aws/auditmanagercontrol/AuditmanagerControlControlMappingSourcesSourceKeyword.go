// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagercontrol


type AuditmanagerControlControlMappingSourcesSourceKeyword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/auditmanager_control#keyword_input_type AuditmanagerControl#keyword_input_type}.
	KeywordInputType *string `field:"required" json:"keywordInputType" yaml:"keywordInputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/auditmanager_control#keyword_value AuditmanagerControl#keyword_value}.
	KeywordValue *string `field:"required" json:"keywordValue" yaml:"keywordValue"`
}

