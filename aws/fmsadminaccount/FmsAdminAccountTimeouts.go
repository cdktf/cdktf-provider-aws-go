// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmsadminaccount


type FmsAdminAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/fms_admin_account#create FmsAdminAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/fms_admin_account#delete FmsAdminAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

