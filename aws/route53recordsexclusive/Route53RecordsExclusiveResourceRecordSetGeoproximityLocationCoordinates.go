// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive


type Route53RecordsExclusiveResourceRecordSetGeoproximityLocationCoordinates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/route53_records_exclusive#latitude Route53RecordsExclusive#latitude}.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/route53_records_exclusive#longitude Route53RecordsExclusive#longitude}.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
}

