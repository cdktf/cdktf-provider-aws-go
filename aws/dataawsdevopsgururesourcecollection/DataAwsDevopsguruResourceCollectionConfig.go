// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdevopsgururesourcecollection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDevopsguruResourceCollectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/devopsguru_resource_collection#type DataAwsDevopsguruResourceCollection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// cloudformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/devopsguru_resource_collection#cloudformation DataAwsDevopsguruResourceCollection#cloudformation}
	Cloudformation interface{} `field:"optional" json:"cloudformation" yaml:"cloudformation"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/devopsguru_resource_collection#tags DataAwsDevopsguruResourceCollection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

