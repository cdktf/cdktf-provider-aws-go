// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#findings_destination DataAwsCloudwatchLogDataProtectionPolicyDocument#findings_destination}
	FindingsDestination *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

