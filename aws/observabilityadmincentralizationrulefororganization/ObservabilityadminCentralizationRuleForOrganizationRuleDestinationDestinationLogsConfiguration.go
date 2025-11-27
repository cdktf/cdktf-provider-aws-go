// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilityadmincentralizationrulefororganization


type ObservabilityadminCentralizationRuleForOrganizationRuleDestinationDestinationLogsConfiguration struct {
	// backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/observabilityadmin_centralization_rule_for_organization#backup_configuration ObservabilityadminCentralizationRuleForOrganization#backup_configuration}
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// logs_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/observabilityadmin_centralization_rule_for_organization#logs_encryption_configuration ObservabilityadminCentralizationRuleForOrganization#logs_encryption_configuration}
	LogsEncryptionConfiguration interface{} `field:"optional" json:"logsEncryptionConfiguration" yaml:"logsEncryptionConfiguration"`
}

