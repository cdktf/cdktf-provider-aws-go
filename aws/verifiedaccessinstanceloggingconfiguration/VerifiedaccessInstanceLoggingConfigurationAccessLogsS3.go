// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessinstanceloggingconfiguration


type VerifiedaccessInstanceLoggingConfigurationAccessLogsS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/verifiedaccess_instance_logging_configuration#enabled VerifiedaccessInstanceLoggingConfiguration#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/verifiedaccess_instance_logging_configuration#bucket_name VerifiedaccessInstanceLoggingConfiguration#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/verifiedaccess_instance_logging_configuration#bucket_owner VerifiedaccessInstanceLoggingConfiguration#bucket_owner}.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/verifiedaccess_instance_logging_configuration#prefix VerifiedaccessInstanceLoggingConfiguration#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

