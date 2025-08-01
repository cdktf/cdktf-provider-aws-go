// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation struct {
	// audit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#audit DataAwsCloudwatchLogDataProtectionPolicyDocument#audit}
	Audit *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit `field:"optional" json:"audit" yaml:"audit"`
	// deidentify block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#deidentify DataAwsCloudwatchLogDataProtectionPolicyDocument#deidentify}
	Deidentify *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify `field:"optional" json:"deidentify" yaml:"deidentify"`
}

