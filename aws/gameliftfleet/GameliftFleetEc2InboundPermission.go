// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gameliftfleet


type GameliftFleetEc2InboundPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/gamelift_fleet#from_port GameliftFleet#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/gamelift_fleet#ip_range GameliftFleet#ip_range}.
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/gamelift_fleet#protocol GameliftFleet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/gamelift_fleet#to_port GameliftFleet#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

