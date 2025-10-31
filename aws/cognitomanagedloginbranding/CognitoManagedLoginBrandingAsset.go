// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitomanagedloginbranding


type CognitoManagedLoginBrandingAsset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/cognito_managed_login_branding#category CognitoManagedLoginBranding#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/cognito_managed_login_branding#color_mode CognitoManagedLoginBranding#color_mode}.
	ColorMode *string `field:"required" json:"colorMode" yaml:"colorMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/cognito_managed_login_branding#extension CognitoManagedLoginBranding#extension}.
	Extension *string `field:"required" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/cognito_managed_login_branding#bytes CognitoManagedLoginBranding#bytes}.
	Bytes *string `field:"optional" json:"bytes" yaml:"bytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/cognito_managed_login_branding#resource_id CognitoManagedLoginBranding#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

