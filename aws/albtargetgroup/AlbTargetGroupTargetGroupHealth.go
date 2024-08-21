// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package albtargetgroup


type AlbTargetGroupTargetGroupHealth struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/alb_target_group#dns_failover AlbTargetGroup#dns_failover}
	DnsFailover *AlbTargetGroupTargetGroupHealthDnsFailover `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/alb_target_group#unhealthy_state_routing AlbTargetGroup#unhealthy_state_routing}
	UnhealthyStateRouting *AlbTargetGroupTargetGroupHealthUnhealthyStateRouting `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

