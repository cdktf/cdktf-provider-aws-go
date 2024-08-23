// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup


type LbTargetGroupTargetGroupHealth struct {
	// dns_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lb_target_group#dns_failover LbTargetGroup#dns_failover}
	DnsFailover *LbTargetGroupTargetGroupHealthDnsFailover `field:"optional" json:"dnsFailover" yaml:"dnsFailover"`
	// unhealthy_state_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lb_target_group#unhealthy_state_routing LbTargetGroup#unhealthy_state_routing}
	UnhealthyStateRouting *LbTargetGroupTargetGroupHealthUnhealthyStateRouting `field:"optional" json:"unhealthyStateRouting" yaml:"unhealthyStateRouting"`
}

