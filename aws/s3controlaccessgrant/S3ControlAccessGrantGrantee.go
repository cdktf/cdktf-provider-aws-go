// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlaccessgrant


type S3ControlAccessGrantGrantee struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/s3control_access_grant#grantee_identifier S3ControlAccessGrant#grantee_identifier}.
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/s3control_access_grant#grantee_type S3ControlAccessGrant#grantee_type}.
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

