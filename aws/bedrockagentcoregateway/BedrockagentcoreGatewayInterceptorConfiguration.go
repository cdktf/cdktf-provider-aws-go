// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregateway


type BedrockagentcoreGatewayInterceptorConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway#interception_points BedrockagentcoreGateway#interception_points}.
	InterceptionPoints *[]*string `field:"required" json:"interceptionPoints" yaml:"interceptionPoints"`
	// input_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway#input_configuration BedrockagentcoreGateway#input_configuration}
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// interceptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_gateway#interceptor BedrockagentcoreGateway#interceptor}
	Interceptor interface{} `field:"optional" json:"interceptor" yaml:"interceptor"`
}

