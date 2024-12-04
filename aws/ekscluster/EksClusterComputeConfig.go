// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterComputeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/eks_cluster#enabled EksCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/eks_cluster#node_pools EksCluster#node_pools}.
	NodePools *[]*string `field:"optional" json:"nodePools" yaml:"nodePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/eks_cluster#node_role_arn EksCluster#node_role_arn}.
	NodeRoleArn *string `field:"optional" json:"nodeRoleArn" yaml:"nodeRoleArn"`
}

