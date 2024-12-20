// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskset


type EcsTaskSetCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/ecs_task_set#capacity_provider EcsTaskSet#capacity_provider}.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/ecs_task_set#weight EcsTaskSet#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/ecs_task_set#base EcsTaskSet#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
}

