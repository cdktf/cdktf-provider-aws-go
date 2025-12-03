// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnconnection


type VpnConnectionTunnel1LogOptionsCloudwatchLogOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#bgp_log_enabled VpnConnection#bgp_log_enabled}.
	BgpLogEnabled interface{} `field:"optional" json:"bgpLogEnabled" yaml:"bgpLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#bgp_log_group_arn VpnConnection#bgp_log_group_arn}.
	BgpLogGroupArn *string `field:"optional" json:"bgpLogGroupArn" yaml:"bgpLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#bgp_log_output_format VpnConnection#bgp_log_output_format}.
	BgpLogOutputFormat *string `field:"optional" json:"bgpLogOutputFormat" yaml:"bgpLogOutputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#log_enabled VpnConnection#log_enabled}.
	LogEnabled interface{} `field:"optional" json:"logEnabled" yaml:"logEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#log_group_arn VpnConnection#log_group_arn}.
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/vpn_connection#log_output_format VpnConnection#log_output_format}.
	LogOutputFormat *string `field:"optional" json:"logOutputFormat" yaml:"logOutputFormat"`
}

