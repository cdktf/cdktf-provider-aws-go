package vpnconnection


type VpnConnectionTunnel2LogOptions struct {
	// cloudwatch_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/vpn_connection#cloudwatch_log_options VpnConnection#cloudwatch_log_options}
	CloudwatchLogOptions *VpnConnectionTunnel2LogOptionsCloudwatchLogOptions `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

