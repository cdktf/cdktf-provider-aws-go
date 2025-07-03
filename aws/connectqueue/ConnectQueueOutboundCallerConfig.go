// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectqueue


type ConnectQueueOutboundCallerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/connect_queue#outbound_caller_id_name ConnectQueue#outbound_caller_id_name}.
	OutboundCallerIdName *string `field:"optional" json:"outboundCallerIdName" yaml:"outboundCallerIdName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/connect_queue#outbound_caller_id_number_id ConnectQueue#outbound_caller_id_number_id}.
	OutboundCallerIdNumberId *string `field:"optional" json:"outboundCallerIdNumberId" yaml:"outboundCallerIdNumberId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/connect_queue#outbound_flow_id ConnectQueue#outbound_flow_id}.
	OutboundFlowId *string `field:"optional" json:"outboundFlowId" yaml:"outboundFlowId"`
}

