// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecGrpcRouteAction struct {
	// weighted_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/appmesh_route#weighted_target AppmeshRoute#weighted_target}
	WeightedTarget interface{} `field:"required" json:"weightedTarget" yaml:"weightedTarget"`
}

