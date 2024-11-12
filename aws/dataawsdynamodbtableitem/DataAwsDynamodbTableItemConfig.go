// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdynamodbtableitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDynamodbTableItemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/dynamodb_table_item#key DataAwsDynamodbTableItem#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/dynamodb_table_item#table_name DataAwsDynamodbTableItem#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/dynamodb_table_item#expression_attribute_names DataAwsDynamodbTableItem#expression_attribute_names}.
	ExpressionAttributeNames *map[string]*string `field:"optional" json:"expressionAttributeNames" yaml:"expressionAttributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/dynamodb_table_item#id DataAwsDynamodbTableItem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/dynamodb_table_item#projection_expression DataAwsDynamodbTableItem#projection_expression}.
	ProjectionExpression *string `field:"optional" json:"projectionExpression" yaml:"projectionExpression"`
}

