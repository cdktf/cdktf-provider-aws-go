// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationStorageServiceConfiguration struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/bedrockagent_flow#s3 BedrockagentFlow#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

