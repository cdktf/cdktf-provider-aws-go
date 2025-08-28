// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamWorkerAccessConfiguration struct {
	// s3_presign block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/sagemaker_workteam#s3_presign SagemakerWorkteam#s3_presign}
	S3Presign *SagemakerWorkteamWorkerAccessConfigurationS3Presign `field:"optional" json:"s3Presign" yaml:"s3Presign"`
}

