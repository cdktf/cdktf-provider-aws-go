// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceVolumeConfigurationManagedEbsVolumeTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ecs_service#resource_type EcsService#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ecs_service#propagate_tags EcsService#propagate_tags}.
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/ecs_service#tags EcsService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

