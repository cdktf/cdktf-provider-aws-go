// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationlftagexpression


type LakeformationLfTagExpressionExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lakeformation_lf_tag_expression#tag_key LakeformationLfTagExpression#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lakeformation_lf_tag_expression#tag_values LakeformationLfTagExpression#tag_values}.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

