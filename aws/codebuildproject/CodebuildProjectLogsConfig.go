// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectLogsConfig struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#cloudwatch_logs CodebuildProject#cloudwatch_logs}
	CloudwatchLogs *CodebuildProjectLogsConfigCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#s3_logs CodebuildProject#s3_logs}
	S3Logs *CodebuildProjectLogsConfigS3Logs `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

