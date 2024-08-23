// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/appflow_flow#bucket_name AppflowFlow#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/appflow_flow#fail_on_first_destination_error AppflowFlow#fail_on_first_destination_error}.
	FailOnFirstDestinationError interface{} `field:"optional" json:"failOnFirstDestinationError" yaml:"failOnFirstDestinationError"`
}

