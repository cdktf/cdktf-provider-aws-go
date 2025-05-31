// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationLogging struct {
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/imagebuilder_infrastructure_configuration#s3_logs ImagebuilderInfrastructureConfiguration#s3_logs}
	S3Logs *ImagebuilderInfrastructureConfigurationLoggingS3Logs `field:"required" json:"s3Logs" yaml:"s3Logs"`
}

