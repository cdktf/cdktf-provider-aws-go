// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile


type ConnectRoutingProfileMediaConcurrencies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/connect_routing_profile#concurrency ConnectRoutingProfile#concurrency}.
	Concurrency *float64 `field:"required" json:"concurrency" yaml:"concurrency"`
	// cross_channel_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/connect_routing_profile#cross_channel_behavior ConnectRoutingProfile#cross_channel_behavior}
	CrossChannelBehavior *ConnectRoutingProfileMediaConcurrenciesCrossChannelBehavior `field:"optional" json:"crossChannelBehavior" yaml:"crossChannelBehavior"`
}

