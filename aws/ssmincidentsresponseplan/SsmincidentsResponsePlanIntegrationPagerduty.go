// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegrationPagerduty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/ssmincidents_response_plan#name SsmincidentsResponsePlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/ssmincidents_response_plan#secret_id SsmincidentsResponsePlan#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/ssmincidents_response_plan#service_id SsmincidentsResponsePlan#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

