// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecs_service#log_driver EcsService#log_driver}.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecs_service#options EcsService#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecs_service#secret_option EcsService#secret_option}
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

