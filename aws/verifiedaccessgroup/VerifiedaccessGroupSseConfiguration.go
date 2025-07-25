// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccessgroup


type VerifiedaccessGroupSseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/verifiedaccess_group#customer_managed_key_enabled VerifiedaccessGroup#customer_managed_key_enabled}.
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/verifiedaccess_group#kms_key_arn VerifiedaccessGroup#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

