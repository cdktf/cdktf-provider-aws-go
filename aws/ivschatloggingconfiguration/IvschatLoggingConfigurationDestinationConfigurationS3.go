// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/ivschat_logging_configuration#bucket_name IvschatLoggingConfiguration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

