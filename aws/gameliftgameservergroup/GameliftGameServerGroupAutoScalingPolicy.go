// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftgameservergroup


type GameliftGameServerGroupAutoScalingPolicy struct {
	// target_tracking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/gamelift_game_server_group#target_tracking_configuration GameliftGameServerGroup#target_tracking_configuration}
	TargetTrackingConfiguration *GameliftGameServerGroupAutoScalingPolicyTargetTrackingConfiguration `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/gamelift_game_server_group#estimated_instance_warmup GameliftGameServerGroup#estimated_instance_warmup}.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}

