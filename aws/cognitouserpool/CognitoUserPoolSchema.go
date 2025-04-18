// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#attribute_data_type CognitoUserPool#attribute_data_type}.
	AttributeDataType *string `field:"required" json:"attributeDataType" yaml:"attributeDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#name CognitoUserPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#developer_only_attribute CognitoUserPool#developer_only_attribute}.
	DeveloperOnlyAttribute interface{} `field:"optional" json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#mutable CognitoUserPool#mutable}.
	Mutable interface{} `field:"optional" json:"mutable" yaml:"mutable"`
	// number_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#number_attribute_constraints CognitoUserPool#number_attribute_constraints}
	NumberAttributeConstraints *CognitoUserPoolSchemaNumberAttributeConstraints `field:"optional" json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#required CognitoUserPool#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// string_attribute_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/cognito_user_pool#string_attribute_constraints CognitoUserPool#string_attribute_constraints}
	StringAttributeConstraints *CognitoUserPoolSchemaStringAttributeConstraints `field:"optional" json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

