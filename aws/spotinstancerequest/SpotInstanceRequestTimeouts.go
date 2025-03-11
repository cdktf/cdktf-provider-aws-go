// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/spot_instance_request#create SpotInstanceRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/spot_instance_request#delete SpotInstanceRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/spot_instance_request#read SpotInstanceRequest#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

