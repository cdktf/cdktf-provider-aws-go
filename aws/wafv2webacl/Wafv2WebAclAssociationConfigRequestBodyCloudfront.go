// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl


type Wafv2WebAclAssociationConfigRequestBodyCloudfront struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/wafv2_web_acl#default_size_inspection_limit Wafv2WebAcl#default_size_inspection_limit}.
	DefaultSizeInspectionLimit *string `field:"required" json:"defaultSizeInspectionLimit" yaml:"defaultSizeInspectionLimit"`
}

