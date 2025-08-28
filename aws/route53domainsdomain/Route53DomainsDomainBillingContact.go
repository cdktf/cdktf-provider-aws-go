// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain


type Route53DomainsDomainBillingContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#address_line_1 Route53DomainsDomain#address_line_1}.
	AddressLine1 *string `field:"optional" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#address_line_2 Route53DomainsDomain#address_line_2}.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#city Route53DomainsDomain#city}.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#contact_type Route53DomainsDomain#contact_type}.
	ContactType *string `field:"optional" json:"contactType" yaml:"contactType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#country_code Route53DomainsDomain#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#email Route53DomainsDomain#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#extra_param Route53DomainsDomain#extra_param}.
	ExtraParam interface{} `field:"optional" json:"extraParam" yaml:"extraParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#fax Route53DomainsDomain#fax}.
	Fax *string `field:"optional" json:"fax" yaml:"fax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#first_name Route53DomainsDomain#first_name}.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#last_name Route53DomainsDomain#last_name}.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#organization_name Route53DomainsDomain#organization_name}.
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#phone_number Route53DomainsDomain#phone_number}.
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#state Route53DomainsDomain#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#zip_code Route53DomainsDomain#zip_code}.
	ZipCode *string `field:"optional" json:"zipCode" yaml:"zipCode"`
}

