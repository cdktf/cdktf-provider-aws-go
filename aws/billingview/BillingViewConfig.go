// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#name BillingView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// data_filter_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#data_filter_expression BillingView#data_filter_expression}
	DataFilterExpression interface{} `field:"optional" json:"dataFilterExpression" yaml:"dataFilterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#description BillingView#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#source_views BillingView#source_views}.
	SourceViews *[]*string `field:"optional" json:"sourceViews" yaml:"sourceViews"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#tags BillingView#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/billing_view#timeouts BillingView#timeouts}
	Timeouts *BillingViewTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

