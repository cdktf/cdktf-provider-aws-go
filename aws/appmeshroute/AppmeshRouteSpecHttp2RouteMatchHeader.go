// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttp2RouteMatchHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_route#name AppmeshRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_route#invert AppmeshRoute#invert}.
	Invert interface{} `field:"optional" json:"invert" yaml:"invert"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttp2RouteMatchHeaderMatch `field:"optional" json:"match" yaml:"match"`
}

