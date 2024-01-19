// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttp2RouteActionWeightedTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#virtual_node AppmeshRoute#virtual_node}.
	VirtualNode *string `field:"required" json:"virtualNode" yaml:"virtualNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#weight AppmeshRoute#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_route#port AppmeshRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

