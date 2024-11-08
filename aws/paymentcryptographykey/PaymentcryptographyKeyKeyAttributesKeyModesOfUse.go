// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykey


type PaymentcryptographyKeyKeyAttributesKeyModesOfUse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#decrypt PaymentcryptographyKey#decrypt}.
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#derive_key PaymentcryptographyKey#derive_key}.
	DeriveKey interface{} `field:"optional" json:"deriveKey" yaml:"deriveKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#encrypt PaymentcryptographyKey#encrypt}.
	Encrypt interface{} `field:"optional" json:"encrypt" yaml:"encrypt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#generate PaymentcryptographyKey#generate}.
	Generate interface{} `field:"optional" json:"generate" yaml:"generate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#no_restrictions PaymentcryptographyKey#no_restrictions}.
	NoRestrictions interface{} `field:"optional" json:"noRestrictions" yaml:"noRestrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#sign PaymentcryptographyKey#sign}.
	Sign interface{} `field:"optional" json:"sign" yaml:"sign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#unwrap PaymentcryptographyKey#unwrap}.
	Unwrap interface{} `field:"optional" json:"unwrap" yaml:"unwrap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#verify PaymentcryptographyKey#verify}.
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/paymentcryptography_key#wrap PaymentcryptographyKey#wrap}.
	Wrap interface{} `field:"optional" json:"wrap" yaml:"wrap"`
}

