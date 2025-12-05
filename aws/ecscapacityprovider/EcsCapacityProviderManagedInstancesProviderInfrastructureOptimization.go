// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInfrastructureOptimization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_capacity_provider#scale_in_after EcsCapacityProvider#scale_in_after}.
	ScaleInAfter *float64 `field:"optional" json:"scaleInAfter" yaml:"scaleInAfter"`
}

