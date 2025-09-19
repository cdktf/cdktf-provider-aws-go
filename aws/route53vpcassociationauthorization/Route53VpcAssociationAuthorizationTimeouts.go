// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53vpcassociationauthorization


type Route53VpcAssociationAuthorizationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/route53_vpc_association_authorization#create Route53VpcAssociationAuthorization#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/route53_vpc_association_authorization#delete Route53VpcAssociationAuthorization#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/route53_vpc_association_authorization#read Route53VpcAssociationAuthorization#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

