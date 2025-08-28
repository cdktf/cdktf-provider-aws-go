// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection


type VpnConnectionTunnel1LogOptions struct {
	// cloudwatch_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/vpn_connection#cloudwatch_log_options VpnConnection#cloudwatch_log_options}
	CloudwatchLogOptions *VpnConnectionTunnel1LogOptionsCloudwatchLogOptions `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

