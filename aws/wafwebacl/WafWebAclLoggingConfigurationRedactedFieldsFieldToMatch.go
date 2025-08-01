// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafwebacl


type WafWebAclLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/waf_web_acl#type WafWebAcl#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/waf_web_acl#data WafWebAcl#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

