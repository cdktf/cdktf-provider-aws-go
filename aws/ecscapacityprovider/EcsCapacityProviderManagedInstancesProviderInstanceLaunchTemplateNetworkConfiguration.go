// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#subnets EcsCapacityProvider#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#security_groups EcsCapacityProvider#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

