// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalwebacl


type WafregionalWebAclLoggingConfigurationRedactedFields struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_web_acl#field_to_match WafregionalWebAcl#field_to_match}
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
}

