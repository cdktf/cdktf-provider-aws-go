// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetRuntimeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#game_session_activation_timeout_seconds GameliftFleet#game_session_activation_timeout_seconds}.
	GameSessionActivationTimeoutSeconds *float64 `field:"optional" json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#max_concurrent_game_session_activations GameliftFleet#max_concurrent_game_session_activations}.
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// server_process block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_fleet#server_process GameliftFleet#server_process}
	ServerProcess interface{} `field:"optional" json:"serverProcess" yaml:"serverProcess"`
}

