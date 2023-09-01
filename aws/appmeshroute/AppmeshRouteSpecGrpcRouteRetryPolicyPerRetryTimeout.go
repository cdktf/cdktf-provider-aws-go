// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appmesh_route#unit AppmeshRoute#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appmesh_route#value AppmeshRoute#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

