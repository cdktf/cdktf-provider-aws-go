// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftalias


type GameliftAliasRoutingStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/gamelift_alias#type GameliftAlias#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/gamelift_alias#fleet_id GameliftAlias#fleet_id}.
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/gamelift_alias#message GameliftAlias#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

