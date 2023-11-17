// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclRuleActionChallenge struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/wafv2_web_acl#custom_request_handling Wafv2WebAcl#custom_request_handling}
	CustomRequestHandling *Wafv2WebAclRuleActionChallengeCustomRequestHandling `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

