// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclAssociationConfigRequestBody struct {
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/wafv2_web_acl#cloudfront Wafv2WebAcl#cloudfront}
	Cloudfront interface{} `field:"optional" json:"cloudfront" yaml:"cloudfront"`
}

