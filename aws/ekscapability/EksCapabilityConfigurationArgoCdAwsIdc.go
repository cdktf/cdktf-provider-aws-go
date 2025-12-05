// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdAwsIdc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#idc_instance_arn EksCapability#idc_instance_arn}.
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/eks_capability#idc_region EksCapability#idc_region}.
	IdcRegion *string `field:"optional" json:"idcRegion" yaml:"idcRegion"`
}

