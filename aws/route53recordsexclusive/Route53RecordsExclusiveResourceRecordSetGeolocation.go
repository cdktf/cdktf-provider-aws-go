// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive


type Route53RecordsExclusiveResourceRecordSetGeolocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/route53_records_exclusive#continent_code Route53RecordsExclusive#continent_code}.
	ContinentCode *string `field:"optional" json:"continentCode" yaml:"continentCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/route53_records_exclusive#country_code Route53RecordsExclusive#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/route53_records_exclusive#subdivision_code Route53RecordsExclusive#subdivision_code}.
	SubdivisionCode *string `field:"optional" json:"subdivisionCode" yaml:"subdivisionCode"`
}

