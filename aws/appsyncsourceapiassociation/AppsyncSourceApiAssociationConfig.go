// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncsourceapiassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncSourceApiAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#description AppsyncSourceApiAssociation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#merged_api_arn AppsyncSourceApiAssociation#merged_api_arn}.
	MergedApiArn *string `field:"optional" json:"mergedApiArn" yaml:"mergedApiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#merged_api_id AppsyncSourceApiAssociation#merged_api_id}.
	MergedApiId *string `field:"optional" json:"mergedApiId" yaml:"mergedApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#source_api_arn AppsyncSourceApiAssociation#source_api_arn}.
	SourceApiArn *string `field:"optional" json:"sourceApiArn" yaml:"sourceApiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#source_api_association_config AppsyncSourceApiAssociation#source_api_association_config}.
	SourceApiAssociationConfig interface{} `field:"optional" json:"sourceApiAssociationConfig" yaml:"sourceApiAssociationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#source_api_id AppsyncSourceApiAssociation#source_api_id}.
	SourceApiId *string `field:"optional" json:"sourceApiId" yaml:"sourceApiId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appsync_source_api_association#timeouts AppsyncSourceApiAssociation#timeouts}
	Timeouts *AppsyncSourceApiAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

