// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminapplication


type SsoadminApplicationPortalOptionsSignInOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/ssoadmin_application#origin SsoadminApplication#origin}.
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/ssoadmin_application#application_url SsoadminApplication#application_url}.
	ApplicationUrl *string `field:"optional" json:"applicationUrl" yaml:"applicationUrl"`
}

