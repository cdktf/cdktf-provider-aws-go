// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shieldproactiveengagement


type ShieldProactiveEngagementEmergencyContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/shield_proactive_engagement#email_address ShieldProactiveEngagement#email_address}.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/shield_proactive_engagement#contact_notes ShieldProactiveEngagement#contact_notes}.
	ContactNotes *string `field:"optional" json:"contactNotes" yaml:"contactNotes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/shield_proactive_engagement#phone_number ShieldProactiveEngagement#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

