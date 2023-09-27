// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointConnectionLogOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ec2_client_vpn_endpoint#enabled Ec2ClientVpnEndpoint#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ec2_client_vpn_endpoint#cloudwatch_log_group Ec2ClientVpnEndpoint#cloudwatch_log_group}.
	CloudwatchLogGroup *string `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ec2_client_vpn_endpoint#cloudwatch_log_stream Ec2ClientVpnEndpoint#cloudwatch_log_stream}.
	CloudwatchLogStream *string `field:"optional" json:"cloudwatchLogStream" yaml:"cloudwatchLogStream"`
}

