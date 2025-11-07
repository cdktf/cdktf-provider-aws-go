// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/ecs_capacity_provider#infrastructure_role_arn EcsCapacityProvider#infrastructure_role_arn}.
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// instance_launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/ecs_capacity_provider#instance_launch_template EcsCapacityProvider#instance_launch_template}
	InstanceLaunchTemplate *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate `field:"required" json:"instanceLaunchTemplate" yaml:"instanceLaunchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/ecs_capacity_provider#propagate_tags EcsCapacityProvider#propagate_tags}.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}

