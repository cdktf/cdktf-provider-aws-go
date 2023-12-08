// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingscheduledaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppautoscalingScheduledActionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#name AppautoscalingScheduledAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#resource_id AppautoscalingScheduledAction#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#scalable_dimension AppautoscalingScheduledAction#scalable_dimension}.
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// scalable_target_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#scalable_target_action AppautoscalingScheduledAction#scalable_target_action}
	ScalableTargetAction *AppautoscalingScheduledActionScalableTargetAction `field:"required" json:"scalableTargetAction" yaml:"scalableTargetAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#schedule AppautoscalingScheduledAction#schedule}.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#service_namespace AppautoscalingScheduledAction#service_namespace}.
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#end_time AppautoscalingScheduledAction#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#id AppautoscalingScheduledAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#start_time AppautoscalingScheduledAction#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appautoscaling_scheduled_action#timezone AppautoscalingScheduledAction#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

