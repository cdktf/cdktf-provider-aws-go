// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecServiceDiscoveryDns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appmesh_virtual_node#hostname AppmeshVirtualNode#hostname}.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appmesh_virtual_node#ip_preference AppmeshVirtualNode#ip_preference}.
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/appmesh_virtual_node#response_type AppmeshVirtualNode#response_type}.
	ResponseType *string `field:"optional" json:"responseType" yaml:"responseType"`
}

