// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_capacity_provider#ec2_instance_profile_arn EcsCapacityProvider#ec2_instance_profile_arn}.
	Ec2InstanceProfileArn *string `field:"required" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_capacity_provider#network_configuration EcsCapacityProvider#network_configuration}
	NetworkConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_capacity_provider#instance_requirements EcsCapacityProvider#instance_requirements}
	InstanceRequirements *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_capacity_provider#monitoring EcsCapacityProvider#monitoring}.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ecs_capacity_provider#storage_configuration EcsCapacityProvider#storage_configuration}
	StorageConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

