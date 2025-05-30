// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttp2RouteMatchQueryParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/appmesh_route#name AppmeshRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttp2RouteMatchQueryParameterMatch `field:"optional" json:"match" yaml:"match"`
}

