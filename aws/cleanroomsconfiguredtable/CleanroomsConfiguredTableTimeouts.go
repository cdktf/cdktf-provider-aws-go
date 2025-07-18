// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cleanrooms_configured_table#create CleanroomsConfiguredTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cleanrooms_configured_table#delete CleanroomsConfiguredTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/cleanrooms_configured_table#update CleanroomsConfiguredTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

