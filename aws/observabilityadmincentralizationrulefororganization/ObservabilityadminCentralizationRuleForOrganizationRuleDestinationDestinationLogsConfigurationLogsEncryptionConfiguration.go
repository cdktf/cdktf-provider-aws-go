// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#encryption_strategy ObservabilityadminCentralizationRuleForOrganization#encryption_strategy}.
	EncryptionStrategy *string `field:"required" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#encryption_conflict_resolution_strategy ObservabilityadminCentralizationRuleForOrganization#encryption_conflict_resolution_strategy}.
	EncryptionConflictResolutionStrategy *string `field:"optional" json:"encryptionConflictResolutionStrategy" yaml:"encryptionConflictResolutionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/observabilityadmin_centralization_rule_for_organization#kms_key_arn ObservabilityadminCentralizationRuleForOrganization#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

