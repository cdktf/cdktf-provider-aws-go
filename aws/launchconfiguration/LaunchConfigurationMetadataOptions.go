// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchconfiguration


type LaunchConfigurationMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/launch_configuration#http_endpoint LaunchConfiguration#http_endpoint}.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/launch_configuration#http_put_response_hop_limit LaunchConfiguration#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/launch_configuration#http_tokens LaunchConfiguration#http_tokens}.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

