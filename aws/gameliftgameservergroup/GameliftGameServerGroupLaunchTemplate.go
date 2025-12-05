// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/gamelift_game_server_group#id GameliftGameServerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/gamelift_game_server_group#name GameliftGameServerGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/gamelift_game_server_group#version GameliftGameServerGroup#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

