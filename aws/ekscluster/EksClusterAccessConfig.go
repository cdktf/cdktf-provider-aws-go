// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterAccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_cluster#authentication_mode EksCluster#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_cluster#bootstrap_cluster_creator_admin_permissions EksCluster#bootstrap_cluster_creator_admin_permissions}.
	BootstrapClusterCreatorAdminPermissions interface{} `field:"optional" json:"bootstrapClusterCreatorAdminPermissions" yaml:"bootstrapClusterCreatorAdminPermissions"`
}

