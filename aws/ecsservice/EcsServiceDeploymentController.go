// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentController struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/ecs_service#type EcsService#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

