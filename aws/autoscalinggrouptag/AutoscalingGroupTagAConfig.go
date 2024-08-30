// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggrouptag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupTagAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/autoscaling_group_tag#autoscaling_group_name AutoscalingGroupTagA#autoscaling_group_name}.
	AutoscalingGroupName *string `field:"required" json:"autoscalingGroupName" yaml:"autoscalingGroupName"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/autoscaling_group_tag#tag AutoscalingGroupTagA#tag}
	Tag *AutoscalingGroupTagTag `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/autoscaling_group_tag#id AutoscalingGroupTagA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

