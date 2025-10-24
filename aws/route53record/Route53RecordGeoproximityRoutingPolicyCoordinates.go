// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53record


type Route53RecordGeoproximityRoutingPolicyCoordinates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/route53_record#latitude Route53Record#latitude}.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/route53_record#longitude Route53Record#longitude}.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
}

