// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#cloudwatch_logs DataAwsCloudwatchLogDataProtectionPolicyDocument#cloudwatch_logs}
	CloudwatchLogs *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#firehose DataAwsCloudwatchLogDataProtectionPolicyDocument#firehose}
	Firehose *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#s3 DataAwsCloudwatchLogDataProtectionPolicyDocument#s3}
	S3 *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestinationS3 `field:"optional" json:"s3" yaml:"s3"`
}

