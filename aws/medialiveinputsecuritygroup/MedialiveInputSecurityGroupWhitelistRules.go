// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveinputsecuritygroup


type MedialiveInputSecurityGroupWhitelistRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/medialive_input_security_group#cidr MedialiveInputSecurityGroup#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

