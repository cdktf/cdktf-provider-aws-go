// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerConnectionPoolHttp2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/appmesh_virtual_node#max_requests AppmeshVirtualNode#max_requests}.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

