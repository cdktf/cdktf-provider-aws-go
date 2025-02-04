// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy


type ResiliencehubResiliencyPolicyPolicy struct {
	// az block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/resiliencehub_resiliency_policy#az ResiliencehubResiliencyPolicy#az}
	Az *ResiliencehubResiliencyPolicyPolicyAz `field:"required" json:"az" yaml:"az"`
	// hardware block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/resiliencehub_resiliency_policy#hardware ResiliencehubResiliencyPolicy#hardware}
	Hardware *ResiliencehubResiliencyPolicyPolicyHardware `field:"required" json:"hardware" yaml:"hardware"`
	// software block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/resiliencehub_resiliency_policy#software ResiliencehubResiliencyPolicy#software}
	SoftwareAttribute *ResiliencehubResiliencyPolicyPolicySoftware `field:"required" json:"softwareAttribute" yaml:"softwareAttribute"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/resiliencehub_resiliency_policy#region ResiliencehubResiliencyPolicy#region}
	Region *ResiliencehubResiliencyPolicyPolicyRegion `field:"optional" json:"region" yaml:"region"`
}

