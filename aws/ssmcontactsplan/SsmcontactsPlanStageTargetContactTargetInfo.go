// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsplan


type SsmcontactsPlanStageTargetContactTargetInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/resources/ssmcontacts_plan#is_essential SsmcontactsPlan#is_essential}.
	IsEssential interface{} `field:"required" json:"isEssential" yaml:"isEssential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/resources/ssmcontacts_plan#contact_id SsmcontactsPlan#contact_id}.
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
}

