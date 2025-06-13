// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser


type IdentitystoreUserName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#family_name IdentitystoreUser#family_name}.
	FamilyName *string `field:"required" json:"familyName" yaml:"familyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#given_name IdentitystoreUser#given_name}.
	GivenName *string `field:"required" json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#formatted IdentitystoreUser#formatted}.
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#honorific_prefix IdentitystoreUser#honorific_prefix}.
	HonorificPrefix *string `field:"optional" json:"honorificPrefix" yaml:"honorificPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#honorific_suffix IdentitystoreUser#honorific_suffix}.
	HonorificSuffix *string `field:"optional" json:"honorificSuffix" yaml:"honorificSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/identitystore_user#middle_name IdentitystoreUser#middle_name}.
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
}

