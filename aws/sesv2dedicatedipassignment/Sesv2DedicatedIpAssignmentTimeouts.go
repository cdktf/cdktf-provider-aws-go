// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2dedicatedipassignment


type Sesv2DedicatedIpAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sesv2_dedicated_ip_assignment#create Sesv2DedicatedIpAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/sesv2_dedicated_ip_assignment#delete Sesv2DedicatedIpAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

