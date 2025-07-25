// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy


type ResiliencehubResiliencyPolicyPolicy struct {
	// az block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/resiliencehub_resiliency_policy#az ResiliencehubResiliencyPolicy#az}
	Az interface{} `field:"optional" json:"az" yaml:"az"`
	// hardware block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/resiliencehub_resiliency_policy#hardware ResiliencehubResiliencyPolicy#hardware}
	Hardware interface{} `field:"optional" json:"hardware" yaml:"hardware"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/resiliencehub_resiliency_policy#region ResiliencehubResiliencyPolicy#region}
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// software block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/resiliencehub_resiliency_policy#software ResiliencehubResiliencyPolicy#software}
	SoftwareAttribute interface{} `field:"optional" json:"softwareAttribute" yaml:"softwareAttribute"`
}

