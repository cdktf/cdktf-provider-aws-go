// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationInstanceMetadataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/imagebuilder_infrastructure_configuration#http_put_response_hop_limit ImagebuilderInfrastructureConfiguration#http_put_response_hop_limit}.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/imagebuilder_infrastructure_configuration#http_tokens ImagebuilderInfrastructureConfiguration#http_tokens}.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

