// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountprimarycontact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountPrimaryContactConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#address_line_1 AccountPrimaryContact#address_line_1}.
	AddressLine1 *string `field:"required" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#city AccountPrimaryContact#city}.
	City *string `field:"required" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#country_code AccountPrimaryContact#country_code}.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#full_name AccountPrimaryContact#full_name}.
	FullName *string `field:"required" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#phone_number AccountPrimaryContact#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#postal_code AccountPrimaryContact#postal_code}.
	PostalCode *string `field:"required" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#account_id AccountPrimaryContact#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#address_line_2 AccountPrimaryContact#address_line_2}.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#address_line_3 AccountPrimaryContact#address_line_3}.
	AddressLine3 *string `field:"optional" json:"addressLine3" yaml:"addressLine3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#company_name AccountPrimaryContact#company_name}.
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#district_or_county AccountPrimaryContact#district_or_county}.
	DistrictOrCounty *string `field:"optional" json:"districtOrCounty" yaml:"districtOrCounty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#id AccountPrimaryContact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#state_or_region AccountPrimaryContact#state_or_region}.
	StateOrRegion *string `field:"optional" json:"stateOrRegion" yaml:"stateOrRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/account_primary_contact#website_url AccountPrimaryContact#website_url}.
	WebsiteUrl *string `field:"optional" json:"websiteUrl" yaml:"websiteUrl"`
}

