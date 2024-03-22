// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#domain_name CustomerprofilesProfile#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#account_number CustomerprofilesProfile#account_number}.
	AccountNumber *string `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#additional_information CustomerprofilesProfile#additional_information}.
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#address CustomerprofilesProfile#address}
	Address *CustomerprofilesProfileAddress `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#attributes CustomerprofilesProfile#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// billing_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#billing_address CustomerprofilesProfile#billing_address}
	BillingAddress *CustomerprofilesProfileBillingAddress `field:"optional" json:"billingAddress" yaml:"billingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#birth_date CustomerprofilesProfile#birth_date}.
	BirthDate *string `field:"optional" json:"birthDate" yaml:"birthDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#business_email_address CustomerprofilesProfile#business_email_address}.
	BusinessEmailAddress *string `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#business_name CustomerprofilesProfile#business_name}.
	BusinessName *string `field:"optional" json:"businessName" yaml:"businessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#business_phone_number CustomerprofilesProfile#business_phone_number}.
	BusinessPhoneNumber *string `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#email_address CustomerprofilesProfile#email_address}.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#first_name CustomerprofilesProfile#first_name}.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#gender_string CustomerprofilesProfile#gender_string}.
	GenderString *string `field:"optional" json:"genderString" yaml:"genderString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#home_phone_number CustomerprofilesProfile#home_phone_number}.
	HomePhoneNumber *string `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#id CustomerprofilesProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#last_name CustomerprofilesProfile#last_name}.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// mailing_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#mailing_address CustomerprofilesProfile#mailing_address}
	MailingAddress *CustomerprofilesProfileMailingAddress `field:"optional" json:"mailingAddress" yaml:"mailingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#middle_name CustomerprofilesProfile#middle_name}.
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#mobile_phone_number CustomerprofilesProfile#mobile_phone_number}.
	MobilePhoneNumber *string `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#party_type_string CustomerprofilesProfile#party_type_string}.
	PartyTypeString *string `field:"optional" json:"partyTypeString" yaml:"partyTypeString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#personal_email_address CustomerprofilesProfile#personal_email_address}.
	PersonalEmailAddress *string `field:"optional" json:"personalEmailAddress" yaml:"personalEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#phone_number CustomerprofilesProfile#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// shipping_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/customerprofiles_profile#shipping_address CustomerprofilesProfile#shipping_address}
	ShippingAddress *CustomerprofilesProfileShippingAddress `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
}

