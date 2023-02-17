package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAudit struct {
	// findings_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudwatch_log_data_protection_policy_document#findings_destination DataAwsCloudwatchLogDataProtectionPolicyDocument#findings_destination}
	FindingsDestination *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationAuditFindingsDestination `field:"required" json:"findingsDestination" yaml:"findingsDestination"`
}

