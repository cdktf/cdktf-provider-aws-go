// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshmesh


type AppmeshMeshSpecServiceDiscovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/appmesh_mesh#ip_preference AppmeshMesh#ip_preference}.
	IpPreference *string `field:"optional" json:"ipPreference" yaml:"ipPreference"`
}

