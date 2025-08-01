// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_attachment#autoscaling_group_name AutoscalingAttachment#autoscaling_group_name}.
	AutoscalingGroupName *string `field:"required" json:"autoscalingGroupName" yaml:"autoscalingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_attachment#elb AutoscalingAttachment#elb}.
	Elb *string `field:"optional" json:"elb" yaml:"elb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_attachment#id AutoscalingAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_attachment#lb_target_group_arn AutoscalingAttachment#lb_target_group_arn}.
	LbTargetGroupArn *string `field:"optional" json:"lbTargetGroupArn" yaml:"lbTargetGroupArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_attachment#region AutoscalingAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

