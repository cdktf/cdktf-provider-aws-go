package ec2


type LaunchTemplateMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#http_endpoint LaunchTemplate#http_endpoint}.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#http_protocol_ipv6 LaunchTemplate#http_protocol_ipv6}.
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#http_put_response_hop_limit LaunchTemplate#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#http_tokens LaunchTemplate#http_tokens}.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/launch_template#instance_metadata_tags LaunchTemplate#instance_metadata_tags}.
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

