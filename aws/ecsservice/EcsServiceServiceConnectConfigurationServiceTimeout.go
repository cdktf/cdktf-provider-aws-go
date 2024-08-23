// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServiceTimeout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/ecs_service#idle_timeout_seconds EcsService#idle_timeout_seconds}.
	IdleTimeoutSeconds *float64 `field:"optional" json:"idleTimeoutSeconds" yaml:"idleTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/ecs_service#per_request_timeout_seconds EcsService#per_request_timeout_seconds}.
	PerRequestTimeoutSeconds *float64 `field:"optional" json:"perRequestTimeoutSeconds" yaml:"perRequestTimeoutSeconds"`
}

