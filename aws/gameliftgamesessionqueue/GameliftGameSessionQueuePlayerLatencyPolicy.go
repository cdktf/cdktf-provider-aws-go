// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftgamesessionqueue


type GameliftGameSessionQueuePlayerLatencyPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_game_session_queue#maximum_individual_player_latency_milliseconds GameliftGameSessionQueue#maximum_individual_player_latency_milliseconds}.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `field:"required" json:"maximumIndividualPlayerLatencyMilliseconds" yaml:"maximumIndividualPlayerLatencyMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/gamelift_game_session_queue#policy_duration_seconds GameliftGameSessionQueue#policy_duration_seconds}.
	PolicyDurationSeconds *float64 `field:"optional" json:"policyDurationSeconds" yaml:"policyDurationSeconds"`
}

