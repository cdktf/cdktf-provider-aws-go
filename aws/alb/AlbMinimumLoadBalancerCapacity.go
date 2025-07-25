// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alb


type AlbMinimumLoadBalancerCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/alb#capacity_units Alb#capacity_units}.
	CapacityUnits *float64 `field:"required" json:"capacityUnits" yaml:"capacityUnits"`
}

