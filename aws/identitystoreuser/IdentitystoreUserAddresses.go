// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#country IdentitystoreUser#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#formatted IdentitystoreUser#formatted}.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#locality IdentitystoreUser#locality}.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#postal_code IdentitystoreUser#postal_code}.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#primary IdentitystoreUser#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#region IdentitystoreUser#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#street_address IdentitystoreUser#street_address}.
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/identitystore_user#type IdentitystoreUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

