// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudhsmv2hsm


type CloudhsmV2HsmTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/cloudhsm_v2_hsm#create CloudhsmV2Hsm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/cloudhsm_v2_hsm#delete CloudhsmV2Hsm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

