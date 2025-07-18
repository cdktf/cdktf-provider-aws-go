// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentConfigurationCustomDataIdentifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#name DataAwsCloudwatchLogDataProtectionPolicyDocument#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#regex DataAwsCloudwatchLogDataProtectionPolicyDocument#regex}.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

