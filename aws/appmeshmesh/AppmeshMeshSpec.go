// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshmesh


type AppmeshMeshSpec struct {
	// egress_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appmesh_mesh#egress_filter AppmeshMesh#egress_filter}
	EgressFilter *AppmeshMeshSpecEgressFilter `field:"optional" json:"egressFilter" yaml:"egressFilter"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/appmesh_mesh#service_discovery AppmeshMesh#service_discovery}
	ServiceDiscovery *AppmeshMeshSpecServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

