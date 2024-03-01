// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalwebacl


type WafregionalWebAclLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/wafregional_web_acl#log_destination WafregionalWebAcl#log_destination}.
	LogDestination *string `field:"required" json:"logDestination" yaml:"logDestination"`
	// redacted_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/wafregional_web_acl#redacted_fields WafregionalWebAcl#redacted_fields}
	RedactedFields *WafregionalWebAclLoggingConfigurationRedactedFields `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

