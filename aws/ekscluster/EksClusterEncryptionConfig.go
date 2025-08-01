// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterEncryptionConfig struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_cluster#provider EksCluster#provider}
	Provider *EksClusterEncryptionConfigProvider `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/eks_cluster#resources EksCluster#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

