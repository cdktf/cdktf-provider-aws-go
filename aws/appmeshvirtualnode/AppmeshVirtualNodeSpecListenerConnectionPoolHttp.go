// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerConnectionPoolHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/appmesh_virtual_node#max_connections AppmeshVirtualNode#max_connections}.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/appmesh_virtual_node#max_pending_requests AppmeshVirtualNode#max_pending_requests}.
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

