// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/kendra_index#create KendraIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/kendra_index#delete KendraIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/kendra_index#update KendraIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

