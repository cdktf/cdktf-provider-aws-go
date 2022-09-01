package vpc


type VpcPeeringConnectionOptionsAccepter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_peering_connection_options#allow_classic_link_to_remote_vpc VpcPeeringConnectionOptions#allow_classic_link_to_remote_vpc}.
	AllowClassicLinkToRemoteVpc interface{} `field:"optional" json:"allowClassicLinkToRemoteVpc" yaml:"allowClassicLinkToRemoteVpc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_peering_connection_options#allow_remote_vpc_dns_resolution VpcPeeringConnectionOptions#allow_remote_vpc_dns_resolution}.
	AllowRemoteVpcDnsResolution interface{} `field:"optional" json:"allowRemoteVpcDnsResolution" yaml:"allowRemoteVpcDnsResolution"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_peering_connection_options#allow_vpc_to_remote_classic_link VpcPeeringConnectionOptions#allow_vpc_to_remote_classic_link}.
	AllowVpcToRemoteClassicLink interface{} `field:"optional" json:"allowVpcToRemoteClassicLink" yaml:"allowVpcToRemoteClassicLink"`
}

