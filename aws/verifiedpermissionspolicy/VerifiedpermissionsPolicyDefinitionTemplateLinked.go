// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy


type VerifiedpermissionsPolicyDefinitionTemplateLinked struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/verifiedpermissions_policy#policy_template_id VerifiedpermissionsPolicy#policy_template_id}.
	PolicyTemplateId *string `field:"required" json:"policyTemplateId" yaml:"policyTemplateId"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/verifiedpermissions_policy#principal VerifiedpermissionsPolicy#principal}
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/verifiedpermissions_policy#resource VerifiedpermissionsPolicy#resource}
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}

