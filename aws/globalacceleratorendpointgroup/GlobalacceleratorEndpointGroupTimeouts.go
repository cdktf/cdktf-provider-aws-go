// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/globalaccelerator_endpoint_group#create GlobalacceleratorEndpointGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/globalaccelerator_endpoint_group#delete GlobalacceleratorEndpointGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/globalaccelerator_endpoint_group#update GlobalacceleratorEndpointGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

