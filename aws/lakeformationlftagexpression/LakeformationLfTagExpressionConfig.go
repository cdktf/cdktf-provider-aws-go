// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationlftagexpression

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LakeformationLfTagExpressionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the LF-Tag Expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lakeformation_lf_tag_expression#name LakeformationLfTagExpression#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lakeformation_lf_tag_expression#catalog_id LakeformationLfTagExpression#catalog_id}
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// A description of the LF-Tag Expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lakeformation_lf_tag_expression#description LakeformationLfTagExpression#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lakeformation_lf_tag_expression#expression LakeformationLfTagExpression#expression}
	Expression interface{} `field:"optional" json:"expression" yaml:"expression"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/lakeformation_lf_tag_expression#region LakeformationLfTagExpression#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

