// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterZonalShiftConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/eks_cluster#enabled EksCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

