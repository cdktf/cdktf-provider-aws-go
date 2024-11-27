// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/ssm_maintenance_window_task#document_version SsmMaintenanceWindowTask#document_version}.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/ssm_maintenance_window_task#parameter SsmMaintenanceWindowTask#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

