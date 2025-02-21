// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource struct {
	// nlb_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/route53recoveryreadiness_resource_set#nlb_resource Route53RecoveryreadinessResourceSet#nlb_resource}
	NlbResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// r53_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/route53recoveryreadiness_resource_set#r53_resource Route53RecoveryreadinessResourceSet#r53_resource}
	R53Resource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

