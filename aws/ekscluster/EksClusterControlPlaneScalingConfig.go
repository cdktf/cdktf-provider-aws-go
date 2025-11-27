// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterControlPlaneScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/eks_cluster#tier EksCluster#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

