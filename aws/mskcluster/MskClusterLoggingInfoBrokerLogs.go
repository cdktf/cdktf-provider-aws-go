// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterLoggingInfoBrokerLogs struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/msk_cluster#cloudwatch_logs MskCluster#cloudwatch_logs}
	CloudwatchLogs *MskClusterLoggingInfoBrokerLogsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/msk_cluster#firehose MskCluster#firehose}
	Firehose *MskClusterLoggingInfoBrokerLogsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/msk_cluster#s3 MskCluster#s3}
	S3 *MskClusterLoggingInfoBrokerLogsS3 `field:"optional" json:"s3" yaml:"s3"`
}

