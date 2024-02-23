// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclloggingconfiguration


type Wafv2WebAclLoggingConfigurationLoggingFilterFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/wafv2_web_acl_logging_configuration#behavior Wafv2WebAclLoggingConfiguration#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/wafv2_web_acl_logging_configuration#condition Wafv2WebAclLoggingConfiguration#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/wafv2_web_acl_logging_configuration#requirement Wafv2WebAclLoggingConfiguration#requirement}.
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
}

