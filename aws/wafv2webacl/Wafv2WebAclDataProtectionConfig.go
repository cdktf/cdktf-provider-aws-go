// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclDataProtectionConfig struct {
	// data_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/wafv2_web_acl#data_protection Wafv2WebAcl#data_protection}
	DataProtection interface{} `field:"optional" json:"dataProtection" yaml:"dataProtection"`
}

