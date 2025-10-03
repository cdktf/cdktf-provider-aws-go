// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53recoveryreadiness_resource_set#domain_name Route53RecoveryreadinessResourceSet#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53recoveryreadiness_resource_set#hosted_zone_arn Route53RecoveryreadinessResourceSet#hosted_zone_arn}.
	HostedZoneArn *string `field:"optional" json:"hostedZoneArn" yaml:"hostedZoneArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53recoveryreadiness_resource_set#record_set_id Route53RecoveryreadinessResourceSet#record_set_id}.
	RecordSetId *string `field:"optional" json:"recordSetId" yaml:"recordSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53recoveryreadiness_resource_set#record_type Route53RecoveryreadinessResourceSet#record_type}.
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
	// target_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/route53recoveryreadiness_resource_set#target_resource Route53RecoveryreadinessResourceSet#target_resource}
	TargetResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource `field:"optional" json:"targetResource" yaml:"targetResource"`
}

