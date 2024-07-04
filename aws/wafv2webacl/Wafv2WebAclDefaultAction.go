// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclDefaultAction struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclDefaultActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclDefaultActionBlock `field:"optional" json:"block" yaml:"block"`
}

