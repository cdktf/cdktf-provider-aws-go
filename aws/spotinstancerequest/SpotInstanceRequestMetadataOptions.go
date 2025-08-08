// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/spot_instance_request#http_endpoint SpotInstanceRequest#http_endpoint}.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/spot_instance_request#http_protocol_ipv6 SpotInstanceRequest#http_protocol_ipv6}.
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/spot_instance_request#http_put_response_hop_limit SpotInstanceRequest#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/spot_instance_request#http_tokens SpotInstanceRequest#http_tokens}.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/spot_instance_request#instance_metadata_tags SpotInstanceRequest#instance_metadata_tags}.
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

