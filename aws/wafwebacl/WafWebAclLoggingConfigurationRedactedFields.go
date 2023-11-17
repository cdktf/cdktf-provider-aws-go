// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafwebacl


type WafWebAclLoggingConfigurationRedactedFields struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/waf_web_acl#field_to_match WafWebAcl#field_to_match}
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
}

