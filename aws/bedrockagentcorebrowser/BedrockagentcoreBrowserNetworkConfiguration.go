// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowser


type BedrockagentcoreBrowserNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagentcore_browser#network_mode BedrockagentcoreBrowser#network_mode}.
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagentcore_browser#vpc_config BedrockagentcoreBrowser#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

