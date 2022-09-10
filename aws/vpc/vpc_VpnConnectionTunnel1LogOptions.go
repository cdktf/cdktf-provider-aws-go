package vpc


type VpnConnectionTunnel1LogOptions struct {
	// cloudwatch_log_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpn_connection#cloudwatch_log_options VpnConnection#cloudwatch_log_options}
	CloudwatchLogOptions *VpnConnectionTunnel1LogOptionsCloudwatchLogOptions `field:"optional" json:"cloudwatchLogOptions" yaml:"cloudwatchLogOptions"`
}

