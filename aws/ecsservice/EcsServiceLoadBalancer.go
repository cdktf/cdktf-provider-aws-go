// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ecs_service#container_name EcsService#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ecs_service#container_port EcsService#container_port}.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// advanced_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ecs_service#advanced_configuration EcsService#advanced_configuration}
	AdvancedConfiguration *EcsServiceLoadBalancerAdvancedConfiguration `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ecs_service#elb_name EcsService#elb_name}.
	ElbName *string `field:"optional" json:"elbName" yaml:"elbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/ecs_service#target_group_arn EcsService#target_group_arn}.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

