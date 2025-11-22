// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskset


type EcsTaskSetScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ecs_task_set#unit EcsTaskSet#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/ecs_task_set#value EcsTaskSet#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

