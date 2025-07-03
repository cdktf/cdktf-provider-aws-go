// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup


type BedrockagentAgentActionGroupApiSchemaS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/bedrockagent_agent_action_group#s3_bucket_name BedrockagentAgentActionGroup#s3_bucket_name}.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/bedrockagent_agent_action_group#s3_object_key BedrockagentAgentActionGroup#s3_object_key}.
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

