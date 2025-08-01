// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resiliencehubresiliencypolicy


type ResiliencehubResiliencyPolicyPolicyHardware struct {
	// Recovery Point Objective (RPO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resiliencehub_resiliency_policy#rpo ResiliencehubResiliencyPolicy#rpo}
	Rpo *string `field:"required" json:"rpo" yaml:"rpo"`
	// Recovery Time Objective (RTO) as a Go duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/resiliencehub_resiliency_policy#rto ResiliencehubResiliencyPolicy#rto}
	Rto *string `field:"required" json:"rto" yaml:"rto"`
}

