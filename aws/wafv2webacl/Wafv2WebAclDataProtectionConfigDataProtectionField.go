// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclDataProtectionConfigDataProtectionField struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/wafv2_web_acl#field_type Wafv2WebAcl#field_type}.
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/wafv2_web_acl#field_keys Wafv2WebAcl#field_keys}.
	FieldKeys *[]*string `field:"optional" json:"fieldKeys" yaml:"fieldKeys"`
}

