// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectroutingprofile


type ConnectRoutingProfileQueueConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/connect_routing_profile#delay ConnectRoutingProfile#delay}.
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/connect_routing_profile#priority ConnectRoutingProfile#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/connect_routing_profile#queue_id ConnectRoutingProfile#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

