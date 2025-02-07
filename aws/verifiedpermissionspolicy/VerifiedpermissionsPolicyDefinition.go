// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinition struct {
	// static block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/verifiedpermissions_policy#static VerifiedpermissionsPolicy#static}
	Static interface{} `field:"optional" json:"static" yaml:"static"`
	// template_linked block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/verifiedpermissions_policy#template_linked VerifiedpermissionsPolicy#template_linked}
	TemplateLinked interface{} `field:"optional" json:"templateLinked" yaml:"templateLinked"`
}

