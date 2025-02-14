// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscloudwatchlogdataprotectionpolicydocument


type DataAwsCloudwatchLogDataProtectionPolicyDocumentStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#data_identifiers DataAwsCloudwatchLogDataProtectionPolicyDocument#data_identifiers}.
	DataIdentifiers *[]*string `field:"required" json:"dataIdentifiers" yaml:"dataIdentifiers"`
	// operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#operation DataAwsCloudwatchLogDataProtectionPolicyDocument#operation}
	Operation *DataAwsCloudwatchLogDataProtectionPolicyDocumentStatementOperation `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/data-sources/cloudwatch_log_data_protection_policy_document#sid DataAwsCloudwatchLogDataProtectionPolicyDocument#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

