// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codecatalystsourcerepository


type CodecatalystSourceRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codecatalyst_source_repository#create CodecatalystSourceRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codecatalyst_source_repository#delete CodecatalystSourceRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codecatalyst_source_repository#update CodecatalystSourceRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

