// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsregistereddomain


type Route53DomainsRegisteredDomainTechContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#address_line_1 Route53DomainsRegisteredDomain#address_line_1}.
	AddressLine1 *string `field:"optional" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#address_line_2 Route53DomainsRegisteredDomain#address_line_2}.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#city Route53DomainsRegisteredDomain#city}.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#contact_type Route53DomainsRegisteredDomain#contact_type}.
	ContactType *string `field:"optional" json:"contactType" yaml:"contactType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#country_code Route53DomainsRegisteredDomain#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#email Route53DomainsRegisteredDomain#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#extra_params Route53DomainsRegisteredDomain#extra_params}.
	ExtraParams *map[string]*string `field:"optional" json:"extraParams" yaml:"extraParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#fax Route53DomainsRegisteredDomain#fax}.
	Fax *string `field:"optional" json:"fax" yaml:"fax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#first_name Route53DomainsRegisteredDomain#first_name}.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#last_name Route53DomainsRegisteredDomain#last_name}.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#organization_name Route53DomainsRegisteredDomain#organization_name}.
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#phone_number Route53DomainsRegisteredDomain#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#state Route53DomainsRegisteredDomain#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/route53domains_registered_domain#zip_code Route53DomainsRegisteredDomain#zip_code}.
	ZipCode *string `field:"optional" json:"zipCode" yaml:"zipCode"`
}

