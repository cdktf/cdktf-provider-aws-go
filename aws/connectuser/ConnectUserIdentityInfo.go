// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserIdentityInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user#email ConnectUser#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user#first_name ConnectUser#first_name}.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user#last_name ConnectUser#last_name}.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/connect_user#secondary_email ConnectUser#secondary_email}.
	SecondaryEmail *string `field:"optional" json:"secondaryEmail" yaml:"secondaryEmail"`
}

