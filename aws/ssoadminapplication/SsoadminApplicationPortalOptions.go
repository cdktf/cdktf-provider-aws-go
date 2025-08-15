// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadminapplication


type SsoadminApplicationPortalOptions struct {
	// sign_in_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/ssoadmin_application#sign_in_options SsoadminApplication#sign_in_options}
	SignInOptions interface{} `field:"optional" json:"signInOptions" yaml:"signInOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/ssoadmin_application#visibility SsoadminApplication#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

