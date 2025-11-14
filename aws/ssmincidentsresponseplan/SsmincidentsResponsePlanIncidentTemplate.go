// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanIncidentTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#impact SsmincidentsResponsePlan#impact}.
	Impact *float64 `field:"required" json:"impact" yaml:"impact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#title SsmincidentsResponsePlan#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#dedupe_string SsmincidentsResponsePlan#dedupe_string}.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#incident_tags SsmincidentsResponsePlan#incident_tags}.
	IncidentTags *map[string]*string `field:"optional" json:"incidentTags" yaml:"incidentTags"`
	// notification_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#notification_target SsmincidentsResponsePlan#notification_target}
	NotificationTarget interface{} `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ssmincidents_response_plan#summary SsmincidentsResponsePlan#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

