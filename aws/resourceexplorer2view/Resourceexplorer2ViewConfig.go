// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceexplorer2view

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Resourceexplorer2ViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#name Resourceexplorer2View#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#default_view Resourceexplorer2View#default_view}.
	DefaultView interface{} `field:"optional" json:"defaultView" yaml:"defaultView"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#filters Resourceexplorer2View#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// included_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#included_property Resourceexplorer2View#included_property}
	IncludedProperty interface{} `field:"optional" json:"includedProperty" yaml:"includedProperty"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#region Resourceexplorer2View#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#scope Resourceexplorer2View#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resourceexplorer2_view#tags Resourceexplorer2View#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

