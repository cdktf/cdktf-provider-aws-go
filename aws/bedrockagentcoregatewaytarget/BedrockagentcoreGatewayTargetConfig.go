// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoregatewaytarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentcoreGatewayTargetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#gateway_identifier BedrockagentcoreGatewayTarget#gateway_identifier}.
	GatewayIdentifier *string `field:"required" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#name BedrockagentcoreGatewayTarget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// credential_provider_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#credential_provider_configuration BedrockagentcoreGatewayTarget#credential_provider_configuration}
	CredentialProviderConfiguration interface{} `field:"optional" json:"credentialProviderConfiguration" yaml:"credentialProviderConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#description BedrockagentcoreGatewayTarget#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#region BedrockagentcoreGatewayTarget#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// target_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#target_configuration BedrockagentcoreGatewayTarget#target_configuration}
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagentcore_gateway_target#timeouts BedrockagentcoreGatewayTarget#timeouts}
	Timeouts *BedrockagentcoreGatewayTargetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

