// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerConnectionPoolGrpc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appmesh_virtual_gateway#max_requests AppmeshVirtualGateway#max_requests}.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

