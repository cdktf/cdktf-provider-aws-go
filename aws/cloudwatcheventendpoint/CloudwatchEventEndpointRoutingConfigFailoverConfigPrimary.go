// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventendpoint


type CloudwatchEventEndpointRoutingConfigFailoverConfigPrimary struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/cloudwatch_event_endpoint#health_check CloudwatchEventEndpoint#health_check}.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
}

