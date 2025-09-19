// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2trafficmirrorfilterrule


type Ec2TrafficMirrorFilterRuleDestinationPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_traffic_mirror_filter_rule#from_port Ec2TrafficMirrorFilterRule#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ec2_traffic_mirror_filter_rule#to_port Ec2TrafficMirrorFilterRule#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

