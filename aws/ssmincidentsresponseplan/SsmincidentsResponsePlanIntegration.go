// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIntegration struct {
	// pagerduty block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/ssmincidents_response_plan#pagerduty SsmincidentsResponsePlan#pagerduty}
	Pagerduty interface{} `field:"optional" json:"pagerduty" yaml:"pagerduty"`
}

