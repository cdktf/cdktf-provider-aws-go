// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupAutoScalingPolicyTargetTrackingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/gamelift_game_server_group#target_value GameliftGameServerGroup#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

