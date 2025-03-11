// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetResourceCreationLimitPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/gamelift_fleet#new_game_sessions_per_creator GameliftFleet#new_game_sessions_per_creator}.
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/gamelift_fleet#policy_period_in_minutes GameliftFleet#policy_period_in_minutes}.
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

