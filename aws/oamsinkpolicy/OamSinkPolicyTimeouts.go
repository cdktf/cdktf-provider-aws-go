// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oamsinkpolicy


type OamSinkPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/oam_sink_policy#create OamSinkPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/oam_sink_policy#delete OamSinkPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/oam_sink_policy#update OamSinkPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

