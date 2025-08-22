// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lb


type LbMinimumLoadBalancerCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/lb#capacity_units Lb#capacity_units}.
	CapacityUnits *float64 `field:"required" json:"capacityUnits" yaml:"capacityUnits"`
}

