// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection


type VpnConnectionTunnel1LogOptionsCloudwatchLogOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/vpn_connection#log_enabled VpnConnection#log_enabled}.
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/vpn_connection#log_group_arn VpnConnection#log_group_arn}.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/vpn_connection#log_output_format VpnConnection#log_output_format}.
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

