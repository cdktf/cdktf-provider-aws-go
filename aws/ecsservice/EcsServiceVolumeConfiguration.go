// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVolumeConfiguration struct {
	// managed_ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ecs_service#managed_ebs_volume EcsService#managed_ebs_volume}
	ManagedEbsVolume *EcsServiceVolumeConfigurationManagedEbsVolume `field:"required" json:"managedEbsVolume" yaml:"managedEbsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/ecs_service#name EcsService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

