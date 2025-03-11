// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsidentitystoreuser


type DataAwsIdentitystoreUserAlternateIdentifierUniqueAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/identitystore_user#attribute_path DataAwsIdentitystoreUser#attribute_path}.
	AttributePath *string `field:"required" json:"attributePath" yaml:"attributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/data-sources/identitystore_user#attribute_value DataAwsIdentitystoreUser#attribute_value}.
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

