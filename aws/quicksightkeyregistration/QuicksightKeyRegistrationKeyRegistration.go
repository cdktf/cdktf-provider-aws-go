// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightkeyregistration


type QuicksightKeyRegistrationKeyRegistration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/quicksight_key_registration#key_arn QuicksightKeyRegistration#key_arn}.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/quicksight_key_registration#default_key QuicksightKeyRegistration#default_key}.
	DefaultKey interface{} `field:"optional" json:"defaultKey" yaml:"defaultKey"`
}

