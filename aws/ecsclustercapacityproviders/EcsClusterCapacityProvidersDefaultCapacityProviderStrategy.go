// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsclustercapacityproviders


type EcsClusterCapacityProvidersDefaultCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/ecs_cluster_capacity_providers#capacity_provider EcsClusterCapacityProviders#capacity_provider}.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/ecs_cluster_capacity_providers#base EcsClusterCapacityProviders#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/ecs_cluster_capacity_providers#weight EcsClusterCapacityProviders#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

