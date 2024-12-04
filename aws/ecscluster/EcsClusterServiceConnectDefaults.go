// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterServiceConnectDefaults struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/ecs_cluster#namespace EcsCluster#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

