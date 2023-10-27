// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttp2RouteMatch struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#header AppmeshRoute#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#method AppmeshRoute#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#path AppmeshRoute#path}
	Path *AppmeshRouteSpecHttp2RouteMatchPath `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#port AppmeshRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#prefix AppmeshRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#query_parameter AppmeshRoute#query_parameter}
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_route#scheme AppmeshRoute#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

