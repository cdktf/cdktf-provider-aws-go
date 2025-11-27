// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceServerSideEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagent_data_source#kms_key_arn BedrockagentDataSource#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

