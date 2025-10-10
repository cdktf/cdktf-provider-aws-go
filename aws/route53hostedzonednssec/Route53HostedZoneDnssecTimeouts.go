// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53hostedzonednssec


type Route53HostedZoneDnssecTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/route53_hosted_zone_dnssec#create Route53HostedZoneDnssec#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/route53_hosted_zone_dnssec#delete Route53HostedZoneDnssec#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/route53_hosted_zone_dnssec#update Route53HostedZoneDnssec#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

