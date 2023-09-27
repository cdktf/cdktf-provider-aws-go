// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivschatroom


type IvschatRoomTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ivschat_room#create IvschatRoom#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ivschat_room#delete IvschatRoom#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ivschat_room#update IvschatRoom#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

