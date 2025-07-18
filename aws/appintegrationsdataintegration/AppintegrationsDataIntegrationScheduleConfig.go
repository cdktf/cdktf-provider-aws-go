// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appintegrationsdataintegration


type AppintegrationsDataIntegrationScheduleConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appintegrations_data_integration#first_execution_from AppintegrationsDataIntegration#first_execution_from}.
	FirstExecutionFrom *string `field:"required" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appintegrations_data_integration#object AppintegrationsDataIntegration#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appintegrations_data_integration#schedule_expression AppintegrationsDataIntegration#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

