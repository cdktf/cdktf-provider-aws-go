// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsidentitystoregroup


type DataAwsIdentitystoreGroupAlternateIdentifier struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/identitystore_group#external_id DataAwsIdentitystoreGroup#external_id}
	ExternalId *DataAwsIdentitystoreGroupAlternateIdentifierExternalId `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/identitystore_group#unique_attribute DataAwsIdentitystoreGroup#unique_attribute}
	UniqueAttribute *DataAwsIdentitystoreGroupAlternateIdentifierUniqueAttribute `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

