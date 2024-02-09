// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventendpoint


type CloudwatchEventEndpointRoutingConfig struct {
	// failover_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/cloudwatch_event_endpoint#failover_config CloudwatchEventEndpoint#failover_config}
	FailoverConfig *CloudwatchEventEndpointRoutingConfigFailoverConfig `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

