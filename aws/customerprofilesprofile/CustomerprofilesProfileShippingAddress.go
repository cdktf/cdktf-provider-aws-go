// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesprofile


type CustomerprofilesProfileShippingAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#address_1 CustomerprofilesProfile#address_1}.
	Address1 *string `field:"optional" json:"address1" yaml:"address1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#address_2 CustomerprofilesProfile#address_2}.
	Address2 *string `field:"optional" json:"address2" yaml:"address2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#address_3 CustomerprofilesProfile#address_3}.
	Address3 *string `field:"optional" json:"address3" yaml:"address3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#address_4 CustomerprofilesProfile#address_4}.
	Address4 *string `field:"optional" json:"address4" yaml:"address4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#city CustomerprofilesProfile#city}.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#country CustomerprofilesProfile#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#county CustomerprofilesProfile#county}.
	County *string `field:"optional" json:"county" yaml:"county"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#postal_code CustomerprofilesProfile#postal_code}.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#province CustomerprofilesProfile#province}.
	Province *string `field:"optional" json:"province" yaml:"province"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/customerprofiles_profile#state CustomerprofilesProfile#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

