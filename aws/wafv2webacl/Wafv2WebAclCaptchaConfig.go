// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclCaptchaConfig struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/wafv2_web_acl#immunity_time_property Wafv2WebAcl#immunity_time_property}
	ImmunityTimeProperty *Wafv2WebAclCaptchaConfigImmunityTimeProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

