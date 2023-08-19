package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentify struct {
	// mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/data-sources/cloudwatch_log_data_protection_policy_document#mask_config DataAwsCloudwatchLogDataProtectionPolicyDocument#mask_config}
	MaskConfig *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperationDeidentifyMaskConfig `field:"required" json:"maskConfig" yaml:"maskConfig"`
}

