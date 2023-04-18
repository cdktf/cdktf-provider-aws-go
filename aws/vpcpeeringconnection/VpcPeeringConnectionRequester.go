package vpcpeeringconnection


type VpcPeeringConnectionRequester struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/vpc_peering_connection#allow_classic_link_to_remote_vpc VpcPeeringConnection#allow_classic_link_to_remote_vpc}.
	AllowClassicLinkToRemoteVpc interface{} `field:"optional" json:"allowClassicLinkToRemoteVpc" yaml:"allowClassicLinkToRemoteVpc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/vpc_peering_connection#allow_remote_vpc_dns_resolution VpcPeeringConnection#allow_remote_vpc_dns_resolution}.
	AllowRemoteVpcDnsResolution interface{} `field:"optional" json:"allowRemoteVpcDnsResolution" yaml:"allowRemoteVpcDnsResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/vpc_peering_connection#allow_vpc_to_remote_classic_link VpcPeeringConnection#allow_vpc_to_remote_classic_link}.
	AllowVpcToRemoteClassicLink interface{} `field:"optional" json:"allowVpcToRemoteClassicLink" yaml:"allowVpcToRemoteClassicLink"`
}

