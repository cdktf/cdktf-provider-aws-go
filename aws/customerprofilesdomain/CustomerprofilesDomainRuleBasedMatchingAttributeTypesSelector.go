// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesdomain


type CustomerprofilesDomainRuleBasedMatchingAttributeTypesSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/customerprofiles_domain#attribute_matching_model CustomerprofilesDomain#attribute_matching_model}.
	AttributeMatchingModel *string `field:"required" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/customerprofiles_domain#address CustomerprofilesDomain#address}.
	Address *[]*string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/customerprofiles_domain#email_address CustomerprofilesDomain#email_address}.
	EmailAddress *[]*string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/customerprofiles_domain#phone_number CustomerprofilesDomain#phone_number}.
	PhoneNumber *[]*string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

