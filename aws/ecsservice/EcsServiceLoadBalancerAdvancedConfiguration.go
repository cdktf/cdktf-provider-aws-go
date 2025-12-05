// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceLoadBalancerAdvancedConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#alternate_target_group_arn EcsService#alternate_target_group_arn}.
	AlternateTargetGroupArn *string `field:"required" json:"alternateTargetGroupArn" yaml:"alternateTargetGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#production_listener_rule EcsService#production_listener_rule}.
	ProductionListenerRule *string `field:"required" json:"productionListenerRule" yaml:"productionListenerRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#role_arn EcsService#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_service#test_listener_rule EcsService#test_listener_rule}.
	TestListenerRule *string `field:"optional" json:"testListenerRule" yaml:"testListenerRule"`
}

