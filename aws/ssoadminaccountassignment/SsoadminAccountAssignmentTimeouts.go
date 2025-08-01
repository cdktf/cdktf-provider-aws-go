// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminaccountassignment


type SsoadminAccountAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssoadmin_account_assignment#create SsoadminAccountAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ssoadmin_account_assignment#delete SsoadminAccountAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

