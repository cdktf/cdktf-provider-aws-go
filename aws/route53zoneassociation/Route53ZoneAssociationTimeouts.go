// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53zoneassociation


type Route53ZoneAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/route53_zone_association#create Route53ZoneAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/route53_zone_association#delete Route53ZoneAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

