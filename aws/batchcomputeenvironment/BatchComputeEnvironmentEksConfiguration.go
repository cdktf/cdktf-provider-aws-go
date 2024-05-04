// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentEksConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/batch_compute_environment#eks_cluster_arn BatchComputeEnvironment#eks_cluster_arn}.
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/batch_compute_environment#kubernetes_namespace BatchComputeEnvironment#kubernetes_namespace}.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

