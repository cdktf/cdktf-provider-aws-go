// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclDefaultActionBlock struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/wafv2_web_acl#custom_response Wafv2WebAcl#custom_response}
	CustomResponse *Wafv2WebAclDefaultActionBlockCustomResponse `field:"optional" json:"customResponse" yaml:"customResponse"`
}

