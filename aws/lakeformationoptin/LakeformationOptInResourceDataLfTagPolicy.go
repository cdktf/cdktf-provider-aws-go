// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationoptin


type LakeformationOptInResourceDataLfTagPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lakeformation_opt_in#resource_type LakeformationOptIn#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lakeformation_opt_in#catalog_id LakeformationOptIn#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lakeformation_opt_in#expression LakeformationOptIn#expression}.
	Expression *[]*string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lakeformation_opt_in#expression_name LakeformationOptIn#expression_name}.
	ExpressionName *string `field:"optional" json:"expressionName" yaml:"expressionName"`
}

