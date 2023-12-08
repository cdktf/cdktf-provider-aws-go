// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oamsink


type OamSinkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/oam_sink#create OamSink#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/oam_sink#delete OamSink#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/oam_sink#update OamSink#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

